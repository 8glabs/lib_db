package lib_db

import (
	"github.com/8glabs/lib_db/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

//*models.Nft
func (repo *Repository) GetNftById(nftId uuid.UUID, data interface{}) error {
	err := repo.DB.Preload("Creator").Preload("Owner").Where("id = ?", nftId).First(data).Error
	if err != nil {
		return err
	}
	return nil
}

// func GetNftById(nftId uuid.UUID, db *sql.DB) (*models.Nft, error) {
// 	sqlStatement := "SELECT * FROM nfts WHERE id = $1"

// 	var nft models.Nft

// 	err := db.QueryRow(sqlStatement, nftId).Scan(&nft.Id, &nft.NftCollectionId, &nft.OwnerId, &nft.BuyerId, &nft.SerialId, &nft.ChainType, &nft.TxnStatus, &nft.StarkexAssetId, &nft.StarkexVaultId, &nft.EthereumContractAddress, &nft.EthereumTokenId, &nft.TxnId, &nft.TxnType, &nft.FixedPrice, &nft.FixedPriceCurrency, &nft.OnchainTxnType, &nft.StarkexTxnId, &nft.StripePaymentIntentId)

// 	// Returning ErrNoRows is expected behavior when no rows found, and it's not problemic. Simplely return an empty map.
// 	if err == sql.ErrNoRows {
// 		// TODO(hanhu): Return self-defined error, to not expose sql package outside
// 		// TODO(hanhu): return nil for nft
// 		return &nft, err
// 	}

// 	if err != nil && err != sql.ErrNoRows {
// 		log.Printf("Query nft by id failed %v", err)
// 	}

// 	return &nft, err
// }

//*models.Nft
func (repo *Repository) GetNftByTxnId(txnId uuid.UUID, data interface{}) error {
	err := repo.DB.Preload("NftCollection").Where("txn_id = ?", txnId).Find(data).Error
	if err != nil {
		return err
	}
	return nil
}

// func GetNftByTxnId(txnId uuid.UUID, db *sql.DB) (*models.Nft, error) {
// 	sqlStatement := "SELECT * FROM nfts WHERE txn_id = $1"

// 	var nft models.Nft

// 	err := db.QueryRow(sqlStatement, txnId).Scan(&nft.Id, &nft.NftCollectionId, &nft.OwnerId, &nft.BuyerId, &nft.SerialId, &nft.ChainType, &nft.TxnStatus, &nft.StarkexAssetId, &nft.StarkexVaultId, &nft.EthereumContractAddress, &nft.EthereumTokenId, &nft.TxnId, &nft.TxnType, &nft.FixedPrice, &nft.FixedPriceCurrency, &nft.OnchainTxnType, &nft.StarkexTxnId, &nft.StripePaymentIntentId)

// 	// Returning ErrNoRows is expected behavior when no rows found, and it's not problemic. Simplely return an empty map.
// 	if err == sql.ErrNoRows {
// 		// TODO(hanhu): Return self-defined error, to not expose sql package outside
// 		return &nft, err
// 	}

// 	if err != nil && err != sql.ErrNoRows {
// 		log.Printf("Query nft by txn id failed %v", err)
// 	}
// 	return &nft, err
// }

//*models.Nft
func (repo *Repository) GetNftBySerialId(serialId int, nftCollectionId uuid.UUID, data interface{}) error {
	err := repo.DB.Where("serial_id = ?", serialId).Where("nft_collection_id = ?", nftCollectionId).First(data).Error
	if err != nil {
		return err
	}
	return err
}

func (repo *Repository) UpdateNft(nft interface{}) error {
	// err := repo.DB.Session(&gorm.Session{FullSaveAssociations: false}).Updates(&nft).Error

	// For some reason, if following the doc and do
	// repo.DB.Omit("NftCollection").Omit("Owner").Updates(&nft).Error
	// doesn't work. NftCollection and Owner will still be inserted
	err := repo.DB.Omit("NftCollection", "Owner").Updates(nft).Error
	if err != nil {
		return err
	}
	return nil
}

// func GetTokenInfoOfNft(nftId uuid.UUID, db *sql.DB) (int, string, error) {
// 	sqlStatement := `SELECT t2.loyalty_percentage, t1.media_url
// 					FROM nfts AS t1, nft_collections AS t2 db
// 					WHERE t1.id = $1
// 					AND t1.nft_collection_id = t2.id`

// 	var loyaltyPercentage int
// 	var mediaUrl string

// 	err := db.QueryRow(sqlStatement, nftId).Scan(
// 		&loyaltyPercentage, &mediaUrl)

// 	// Returning ErrNoRows is expected behavior when no rows found, and it's not problemic. Simplely return an empty map.
// 	if err == sql.ErrNoRows {
// 		// TODO(hanhu): Return self-defined error, to not expose sql package outside
// 		return -1, "", err
// 	}

// 	if err != nil && err != sql.ErrNoRows {
// 		log.Printf("Query user by id failed %v", err)
// 	}

// 	return loyaltyPercentage, mediaUrl, err
// }

// func UpdateNftLockStatusInCollection(nftCollectionId uuid.UUID, fromStatus models.TxnAuthority, toStatus models.TxnAuthority, db *sql.DB) error {
// 	sqlStatement := `UPDATE nfts
// 		SET nft_lock_status = $3
// 		WHERE nfts.nft_collection_id = $1
// 		AND nfts.nft_lock_status = $2`

// 	_, err := db.Exec(sqlStatement, nftCollectionId, fromStatus, toStatus)

// 	if err != nil && err != sql.ErrNoRows {
// 		log.Printf("update nft by id failed %v", err)
// 	}

// 	return err
// }

//*[]models.Nft
func (repo *Repository) GetOwnedNfts(userId uint64, data interface{}) error {
	err := repo.DB.Preload(
		"Creator").Preload("Owner").Where("owner_id = ?", userId).Find(data).Error
	if err != nil {
		return err
	}
	return nil
}

//*[]models.Nft
func (repo *Repository) GetOnsaleNftsOfCollection(nftCollectionId uuid.UUID, data interface{}) error {
	err := repo.DB.Preload("NftCollection").Preload("Owner").Where(
		"nft_collection_id = ?", nftCollectionId).Where(
		"txn_status = ?", models.TXN_STATUS_LISTED).Find(data).Error
	if err != nil {
		return nil
	}
	return nil
}
