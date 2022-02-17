package lib_db

import (
	"github.com/8glabs/lib_db/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type NftCollectionDigest struct {
	MediaUrl string `json:"media_url"`
}

//*models.NftCollection
func (repo *Repository) GetNftCollection(nftCollectionId uuid.UUID, data interface{}) error {
	// Get main NFT information
	err := repo.DB.Preload("Nfts").Preload("Nfts.Creator").Where("id = ?", nftCollectionId).First(data).Error
	if err != nil {
		return err
	}
	return nil
}

//*[]models.NftCollection
func (repo *Repository) GetCreatedDrops(userId uint64, from int, to int, idArray []uint64, data interface{}) error {

	if err := repo.DB.Preload("Nfts").Preload("Nfts.Creator").Where(
		"id in (?)", idArray).Where(
		"nft_collections.status !=?", models.SECONDARY_SALE).Offset(from).Limit(to).Find(
		data).Error; err != nil {
		return err
	}
	return nil
}

//*[]models.NftCollection,
func (repo *Repository) GetCreatedMoments(userId uint64, from int, to int, idArray []uint64, data interface{}) error {

	if err := repo.DB.Preload("Nfts").Preload("Nfts.Creator").Where(
		"id in (?)", idArray).Where(
		"nft_collections.status=?", models.SECONDARY_SALE).Offset(from).Limit(to).Find(
		data).Error; err != nil {
		return err
	}
	return nil
}

//*[]models.NftCollection,
func (repo *Repository) GetCreatedNftCollections(userId uint64, from int, to int, idArray []uint64, data interface{}) error {
	if err := repo.DB.Preload("Nfts").Preload("Nfts.Creator").Where(
		"id in (?)", idArray).Offset(from).Limit(to).Find(
		&data).Error; err != nil {
		return err
	}

	return nil
}

// *[]map[string]string,

func (repo *Repository) CreatedNftCollectionDigest(userId int64, idArray []uint64, data interface{}) error {

	if err := repo.DB.Preload("Nfts").Preload("Nfts.Creator").Where(
		"id in (?)", idArray).Find(
		data).Error; err != nil {
		return err
	}
	// digestInfo := make([]map[string]string, len(nftCollections))
	// for i, nftCollection := range nftCollections {
	// 	nfts := *nftCollection.Nfts

	// 	digestInfo[i] = make(map[string]string)
	// 	if len(nfts) != 0 {
	// 		digestInfo[i]["mediaUrl"] = nfts[0].CoverImageUrl //? use index 0？
	// 	} else {
	// 		digestInfo[i]["mediaUrl"] = ""
	// 	}
	// }
	return nil
}

//*[]models.Nft,
func (repo *Repository) GetListedNftsOfCollection(nftCollectionId uuid.UUID, data interface{}) error {
	if err := repo.DB.Where(
		"nft_collection_id = ? AND txn_status IN ?", nftCollectionId, []models.TxnStatus{models.TXN_STATUS_LISTED}).Find(
		data).Error; err != nil {
		return err
	}
	return nil
}

func (repo *Repository) UpdateNftCollection(nftCollection interface{}) error {
	if err := repo.DB.Updates(nftCollection).Error; err != nil {
		return err
	}
	return nil
}

func (repo *Repository) StartSecondarySale(nftCollectionId uuid.UUID, nftCollectionModel, nftModel interface{}) error {
	err := repo.DB.Model(
		nftCollectionModel).Where(
		"id = ?", nftCollectionId).Update(
		"status", models.NftCollectionStatus(models.SECONDARY_SALE)).Error
	if err != nil {
		return err
	}
	err = repo.DB.Model(nftModel).Where(
		"nft_collection_id = ?", nftCollectionId).Where(
		"txn_authority = ?", models.TxnAuthority(models.CREATOR)).Update(
		"txn_authority", models.TxnAuthority(models.OWNER)).Error
	return err
}
