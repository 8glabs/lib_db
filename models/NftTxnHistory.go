package models

import (
	"time"

	"github.com/google/uuid"
)

type NftTxnHistory struct {
	Id               uint64    `json:"id" gorm:"type:bigserial primary key not null" fake:"skip"`
	BuyerId          uint64    `json:"buyer_id" gorm:"type:bigint" fake:"skip"`
	SellerId         uint64    `json:"seller_id" gorm:"type:bigint" fake:"skip"`
	NftId            uuid.UUID `json:"nft_id" gorm:"type:uuid not null" fake:"skip"`
	NftCollectionId  uuid.UUID `json:"nft_collection_id" gorm:"type:uuid not null" fake:"skip"`
	TxnType          TxnType   `json:"txn_type"`
	FixedTxnPrice    float64   `json:"fixed_price" gorm:"type:double precision" fake:"{number:1,20}"`
	TxnPriceCurrency string    `json:"fixed_price_currency" gorm:"type:varchar(50)"`
	TxnChainType     ChainType `json:"txn_chain_type" sql:"type:ChainType default 'UNKNOWN'"`
	SolanaTxnSig     string    `json:"solana_txn_signature" gorm:"type:varchar(100)"`
	TxnCompleteTime  time.Time `json:"txn_complete_time" gorm:"type:timestamp without time zone default current_timestamp" fake:"skip"`
	Nft              *Nft      `json:"nft" gorm:"foreignKey:NftId" fake:"skip"`
	Buyer            *User     `json:"buyer" gorm:"foreignKey:BuyerId" fake:"skip"`
	Seller           *User     `json:"seller" gorm:"foreignKey:SellerId" fake:"skip"`
}
