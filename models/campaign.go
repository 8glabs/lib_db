package models

import "github.com/google/uuid"

type Campaign struct {
	PrimarySalePrice float64        `json:"primary_sale_price" gorm:"type:integer not null" fake:"{number:1,10}"` //add to campaogn
	NftCollectionId  uuid.UUID      `json:"nft_collection_id" gorm:"type:uuid" fake:"skip"`
	CreatorId        uint64         `json:"creator_id" gorm:"type:bigint not null" fake:"skip"`
	NftCollection    *NftCollection `json:"nft_collection" gorm:"foreignKey:NftCollectionId" fake:"skip"`
	Creator          *User          `json:"creator" gorm:"foreignKey:CreatorId" fake:"skip"`
}
