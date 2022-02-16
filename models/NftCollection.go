package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type NftCollectionStatus string

const (
	GIFTING        NftCollectionStatus = "GIFTING"
	PRIMARY_SALE   NftCollectionStatus = "PRIMARY_SALE"
	SECONDARY_SALE NftCollectionStatus = "SECONDARY_SALE"
)

type Rarity string

const (
	Common   Rarity = "Common"
	Uncommon Rarity = "Uncommon"
	Rare     Rarity = "Rare"
	Heroic   Rarity = "Heroic"
	Mythic   Rarity = "Mythic"
)

type ChainType string

const (
	UNKNOWN             ChainType = "UNKNOWN"
	CHAIN_TYPE_ETHEREUM ChainType = "ETHEREUM"
	CHAIN_TYPE_STARKEX  ChainType = "STARKEX"
	CHAIN_TYPE_SOLANA   ChainType = "SOLANA"
	CHAIN_TYPE_LOOPRING ChainType = "LOOPRING"
)

//1 step to change nftCollection 1->n nft
//2 step change NftCollection 1->1 xxxCampaign
type NftCollection struct {
	Id                uuid.UUID           `json:"id" gorm:"type:uuid primary key" fake:"skip"`
	CollectionName    string              `json:"collection_name" gorm:"type:varchar(255) not null" fake:"{animal}"`
	Tags              pq.StringArray      `json:"tags" gorm:"type:varchar(255)[]" fakesize:"3"`
	NftsAmount        int                 `json:"nfts_amount" gorm:"type:integer not null" fake:"{number:1,50}"`
	HighestSalePrice  float64             `json:"highest_sale_price" gorm:"type:double precision" fake:"skip"`
	Description       string              `json:"description" gorm:"type:varchar(1000)" fake:"{sentence:5}"`
	CreatedAt         time.Time           `json:"created_at" gorm:"type:timestamp without time zone default current_timestamp" fake:"skip"`
	Rarity            Rarity              `json:"rarity" sql:"type:Rarity" fake:"skip"`
	Nfts              *[]Nft              `json:"nfts" gorm:"foreignKey:NftCollectionId" fake:"skip"`
	Status            NftCollectionStatus `json:"nft_collection_status" sql:"type:NftCollectionStatus"`      //add to campaogn
	LoyaltyPercentage int                 `json:"loyal_percentage" gorm:"type:integer" fake:"{number:1,10}"` //add to campaogn
}

func (nftCollection *NftCollection) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	nftCollection.Id = id
	return nil
}
