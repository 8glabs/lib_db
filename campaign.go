package lib_db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/8glabs/lib_db/models"
	"github.com/google/uuid"
)

func (repo *Repository) CreateGiftingCampaign(campaign models.GiftingCampaign) error {
	err := repo.DB.Create(&campaign).Error // pass pointer of data to Create
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetGiftingCampaignByCollectionId(nftCollectionId uuid.UUID) (*models.GiftingCampaign, error) {
	var giftingCampaign models.GiftingCampaign
	err := repo.DB.Where("nft_collection_id = ?", nftCollectionId).First(&giftingCampaign).Error
	if err != nil {
		return nil, err
	}
	return &giftingCampaign, nil
}

func (repo *Repository) CreatePrimarySaleCampaign(campaign models.PrimarySaleCampaign) error {
	err := repo.DB.Create(&campaign).Error // pass pointer of data to Create
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetPrimarySaleCampaignByCollectionId(nftCollectionId uuid.UUID) (*models.PrimarySaleCampaign, error) {
	var primarySaleCampaign models.PrimarySaleCampaign
	err := repo.DB.Where("nft_collection_id = ?", nftCollectionId).First(&primarySaleCampaign).Error
	if err != nil {
		return nil, err
	}
	return &primarySaleCampaign, nil
}

func (repo *Repository) GetGiftingCampaignById(giftingCampaignId uuid.UUID) (*models.GiftingCampaign, error) {
	var giftingCampaign models.GiftingCampaign
	err := repo.DB.Preload("NftCollection").Preload("NftCollection.Nfts").First(&giftingCampaign).Error
	if err != nil {
		return nil, err
	}
	fmt.Println("Got nft collection is", giftingCampaign.NftCollection)
	return &giftingCampaign, nil
}

func (repo *Repository) UpdateGiftingCampaign(giftingCampaign *models.GiftingCampaign) error {
	err := repo.DB.Omit("NftCollection", "Creator").Updates(giftingCampaign).Error
	if err != nil {
		return err
	}
	return nil
}
func SendGiftingNft(giftingCampaignId uuid.UUID, receiptId uint64, db *sql.DB) error {
	sqlStatement := `
		UPDATE nfts
		SET owner_id = $1
		FROM nfts as joined
		INNER JOIN gifting_campaigns ON joined.nft_collection_id = gifting_campaigns.nft_collection_id
		WHERE nfts.serial_id = gifting_campaigns.next_available_serial_id
		AND nfts.nft_collection_id = gifting_campaigns.nft_collection_id
		AND gifting_campaigns.id = $2;`

	_, err := db.Exec(sqlStatement, receiptId, giftingCampaignId)

	if err != nil && err != sql.ErrNoRows {
		log.Printf("Set owner of gifting nft failed %v", err)
	}

	if err == sql.ErrNoRows {
		return err
	}

	if err != nil && err != sql.ErrNoRows {
		log.Printf("Query gifting campaign by id failed %v", err)
		return err
	}
	// TODO(https://github.com/8glabs/creatornfts/issues/41): Make it atomic
	// Increment next available gifting id
	sqlStatement = `UPDATE gifting_campaigns
		SET next_available_serial_id = next_available_serial_id + 1
		WHERE gifting_campaigns.id = $1;`
	_, err = db.Exec(sqlStatement, giftingCampaignId)
	if err != nil {
		return err
	}
	return nil
}
