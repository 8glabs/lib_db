package lib_db

import (
	"github.com/google/uuid"
)

//*models.NftTxnHistory
func (repo *Repository) GetHighestHistorySalePrice(nftCollectionId uuid.UUID, data interface{}) error {
	err := repo.DB.Model(data).Select(
		"fixed_txn_price").Where("nft_collection_id = ?", nftCollectionId).Order(
		"fixed_txn_price desc").First(data).Error
	if err != nil {
		return err
	}
	// return highestTxnHistory.FixedTxnPrice, nil
	return nil

}
