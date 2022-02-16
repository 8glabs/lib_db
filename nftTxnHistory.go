package lib_db

import (
	"database/sql"
	"log"

	"github.com/8glabs/lib_db/models"
	"github.com/google/uuid"
)

func AddNftTxnHistory(nftTxnHistory *models.NftTxnHistory, db *sql.DB) (uint64, error) {
	sqlStatement := `INSERT INTO nft_txn_history (buyer_id, seller_id, nft_id, txn_price, txn_price_currency)
					VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := db.QueryRow(sqlStatement, nftTxnHistory.BuyerId, nftTxnHistory.SellerId, nftTxnHistory.NftId, nftTxnHistory.FixedTxnPrice, nftTxnHistory.TxnPriceCurrency).Scan(&nftTxnHistory.Id)

	// Returning ErrNoRows is expected behavior when no rows found, and it's not problemic. Simplely return an empty map.
	if err != nil {
		log.Printf("Insert nft transaction history failed %v", err)
	}
	return nftTxnHistory.Id, err
}

func (repo *Repository) GetHighestHistorySalePrice(nftCollectionId uuid.UUID) (float64, error) {
	var highestTxnHistory models.NftTxnHistory
	err := repo.DB.Model(&models.NftTxnHistory{}).Select(
		"fixed_txn_price").Where("nft_collection_id = ?", nftCollectionId).Order(
		"fixed_txn_price desc").First(&highestTxnHistory).Error
	if err != nil {
		return 0, err
	}
	return highestTxnHistory.FixedTxnPrice, nil
}
