package lib_db

import (
	"fmt"

	"github.com/8glabs/lib_db/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type NftCollectionDigest struct {
	MediaUrl string `json:"media_url"`
}

func (repo *Repository) GetNftCollection(nftCollectionId uuid.UUID) (*models.NftCollection, error) {
	// Get main NFT information
	var nftCollection models.NftCollection
	err := repo.DB.Preload("Nfts").Preload("Nfts.Creator").Where("id = ?", nftCollectionId).First(&nftCollection).Error
	if err != nil {
		return nil, err
	}
	return &nftCollection, nil
}

func (repo *Repository) GetCreatedDrops(userId uint64, from int, to int) (*[]models.NftCollection, error) {

	var nfts []models.Nft
	if err := repo.DB.Where("creator_id = ?", userId).Find(&nfts).Error; err != nil {
		return nil, err
	}

	idMap := make(map[uuid.UUID]int)
	for _, nft := range nfts {
		if _, ok := idMap[nft.Id]; !ok {
			idMap[nft.Id] = 1
		}
	}

	idArray := make([]uuid.UUID, len(idMap))
	for k, _ := range idMap {
		idArray = append(idArray, k)
	}

	var nftCollections []models.NftCollection
	if err := repo.DB.Preload("Nfts").Preload("Nfts.Creator").Where(
		"id in (?)", idArray).Where(
		"nft_collections.status !=?", models.SECONDARY_SALE).Offset(from).Limit(to).Find(
		&nftCollections).Error; err != nil {
		return nil, err
	}
	return &nftCollections, nil
}

func (repo *Repository) GetCreatedMoments(userId uint64, from int, to int) (*[]models.NftCollection, error) {

	var nfts []models.Nft
	if err := repo.DB.Where("creator_id = ?", userId).Find(&nfts).Error; err != nil {
		return nil, err
	}

	idMap := make(map[uuid.UUID]int)
	for _, nft := range nfts {
		if _, ok := idMap[nft.Id]; !ok {
			idMap[nft.Id] = 1
		}
	}

	idArray := make([]uuid.UUID, len(idMap))
	for k, _ := range idMap {
		idArray = append(idArray, k)
	}

	var nftCollections []models.NftCollection
	if err := repo.DB.Preload("Nfts").Preload("Nfts.Creator").Where(
		"id in (?)", idArray).Where(
		"nft_collections.status=?", models.SECONDARY_SALE).Offset(from).Limit(to).Find(
		&nftCollections).Error; err != nil {
		return nil, err
	}
	return &nftCollections, nil
}

func (repo *Repository) GetCreatedNftCollections(userId uint64, from int, to int) (*[]models.NftCollection, error) {
	var nfts []models.Nft
	if err := repo.DB.Where("creator_id = ?", userId).Find(&nfts).Error; err != nil {
		return nil, err
	}

	idMap := make(map[uuid.UUID]int)
	for _, nft := range nfts {
		if _, ok := idMap[nft.Id]; !ok {
			idMap[nft.Id] = 1
		}
	}

	idArray := make([]uuid.UUID, len(idMap))
	for k, _ := range idMap {
		idArray = append(idArray, k)
	}

	var nftCollections []models.NftCollection
	if err := repo.DB.Preload("Nfts").Preload("Nfts.Creator").Where(
		"id in (?)", idArray).Offset(from).Limit(to).Find(
		&nftCollections).Error; err != nil {
		return nil, err
	}
	fmt.Println("Got ", len(nftCollections), " collections")
	return &nftCollections, nil
}

func (repo *Repository) CreatedNftCollectionDigest(userId int64) (*[]map[string]string, error) {
	var nfts []models.Nft
	if err := repo.DB.Where("creator_id = ?", userId).Find(&nfts).Error; err != nil {
		return nil, err
	}

	idMap := make(map[uuid.UUID]int)
	for _, nft := range nfts {
		if _, ok := idMap[nft.Id]; !ok {
			idMap[nft.Id] = 1
		}
	}

	idArray := make([]uuid.UUID, len(idMap))
	for k, _ := range idMap {
		idArray = append(idArray, k)
	}

	var nftCollections []models.NftCollection
	if err := repo.DB.Preload("Nfts").Preload("Nfts.Creator").Where(
		"id in (?)", idArray).Find(
		&nftCollections).Error; err != nil {
		return nil, err
	}
	digestInfo := make([]map[string]string, len(nftCollections))
	for i, nftCollection := range nftCollections {
		nfts := *nftCollection.Nfts

		digestInfo[i] = make(map[string]string)
		if len(nfts) != 0 {
			digestInfo[i]["mediaUrl"] = nfts[0].CoverImageUrl //? use index 0？
		} else {
			digestInfo[i]["mediaUrl"] = ""
		}
	}
	return &digestInfo, nil
}

func (repo *Repository) GetListedNftsOfCollection(nftCollectionId uuid.UUID) (*[]models.Nft, error) {
	var circulatingNfts []models.Nft
	if err := repo.DB.Where(
		"nft_collection_id = ? AND txn_status IN ?", nftCollectionId, []models.TxnStatus{models.TXN_STATUS_LISTED}).Find(
		&circulatingNfts).Error; err != nil {
		return nil, err
	}
	return &circulatingNfts, nil
}

func (repo *Repository) UpdateNftCollection(nftCollection *models.NftCollection) error {
	if err := repo.DB.Updates(nftCollection).Error; err != nil {
		return err
	}
	return nil
}

func (repo *Repository) StartSecondarySale(nftCollectionId uuid.UUID) error {
	err := repo.DB.Model(
		&models.NftCollection{}).Where(
		"id = ?", nftCollectionId).Update(
		"status", models.NftCollectionStatus(models.SECONDARY_SALE)).Error
	if err != nil {
		return err
	}
	err = repo.DB.Model(&models.Nft{}).Where(
		"nft_collection_id = ?", nftCollectionId).Where(
		"txn_authority = ?", models.TxnAuthority(models.CREATOR)).Update(
		"txn_authority", models.TxnAuthority(models.OWNER)).Error
	return err
}
