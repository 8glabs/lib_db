package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TxnType string

const (
	TXN_TYPE_FIXED_PRICE          TxnType = "FIXED_PRICE"
	TXN_TYPE_FIRST_PRICE_AUCTION  TxnType = "FIRST_PRICE_AUCTION"
	TXN_TYPE_SECOND_PRICE_AUCTION TxnType = "SECOND_PRICE_AUCTION"
)

type TxnStatus string

const (
	TXN_STATUS_NOT_LISTED   TxnStatus = "NOT_LISTED"
	TXN_STATUS_LISTED       TxnStatus = "LISTED"
	TXN_STATUS_PENDING      TxnStatus = "PENDING"
	TXN_STATUS_PAYMENT_MADE TxnStatus = "PAYMENT_MADE"
	TXN_STATUS_FAILED       TxnStatus = "FAILED"
)

type TxnAuthority string

const (
	ADMIN   TxnAuthority = "ADMIN"
	PROGRAM TxnAuthority = "PROGRAM"
	CREATOR TxnAuthority = "CREATOR"
	OWNER   TxnAuthority = "OWNER"
)

// nftCollection 1->n nft

type Nft struct {
	Id                         uuid.UUID           `json:"id" gorm:"type:uuid primary key" fake:"skip"`
	NftCollectionId            uuid.UUID           `json:"nft_collection_id" gorm:"type:uuid not null" fake:"skip"`
	OwnerId                    uint64              `json:"owner_id" gorm:"type:bigint" fake:"skip"`
	BuyerId                    uint64              `json:"buyer_id" gorm:"type:bigint" fake:"skip"`
	SerialId                   int                 `json:"serial_id" gorm:"type:integer not null" fake:"skip"`
	TxnStatus                  TxnStatus           `json:"txn_status" sql:"type:TxnStatus"`
	StarkexAssetId             string              `json:"starkex_asset_id" gorm:"type:varchar(255)" fake:"skip"`
	StarkexVaultId             *int64              `json:"starkex_vault_id" gorm:"type:bigint" fake:"skip"`
	EthereumContractAddress    string              `json:"ethereum_contract_address" gorm:"type:varchar(255)" fake:"{bitcoinaddress}"`
	EthereumTokenId            string              `json:"ethereum_token_id" gorm:"type:varchar(255)" fake:"skip"`
	SolanaOwnershipPrivateKey  string              `json:"solana_ownership_private_key" gorm:"type:varchar(255) not null default ''" fake:"skip"`
	TxnId                      uuid.UUID           `json:"txn_id" gorm:"type:uuid default null" fake:"skip"`
	InitialNftCollectionStatus NftCollectionStatus `json:"initial_nft_collection_status" sql:"type:NftCollectionStatus"`
	LastTxnPrice               float64             `json:"last_txn_price" gorm:"type:double precision default 0" fake:"skip"`
	TxnType                    TxnType             `json:"txn_type"`
	FixedPrice                 float64             `json:"fixed_price" gorm:"type:double precision" fake:"{number:1,20}"`
	FixedPriceCurrency         string              `json:"fixed_price_currency" gorm:"type:varchar(50)"`
	StarkexTxnId               string              `json:"starkex_txn_id" gorm:"type:varchar(255) default ''" fake:"skip"`
	StripePaymentIntentId      string              `json:"stripe_payment_intent_id" gorm:"type:varchar(255) default ''"`
	TxnAuthority               TxnAuthority        `json:"txn_authority" sql:"type:TxnAuthority" fake:"skip"`
	NftCollection              *NftCollection      `json:"nft_collection" gorm:"foreignKey:NftCollectionId" fake:"skip"`
	Owner                      *User               `json:"user" gorm:"foreignKey:OwnerId" fake:"skip"`

	MediaType                      MediaType `json:"media_type" sql:"type:MediaType" fake:"skip"`
	MediaUrl                       string    `json:"media_url" gorm:"varchar(255) not null" fake:"{url}"`
	CoverImageUrl                  string    `json:"cover_image_url" gorm:"varchar(255)" fake:"{url}"`
	ChainType                      ChainType `json:"chain_type" sql:"type:ChainType default 'UNKNOWN'"`
	SolanaProgramSecretKey         string    `json:"solana_program_secret_key" gorm:"type:varchar(255) not null" fake:"skip"`
	SolanaMetadataAccountSecretKey string    `json:"solana_metadata_account_secret_key" gorm:"type:varchar(255) not null default ''" fake:"skip"`
	CreatorId                      uint64    `json:"creator_id"  gorm:"type:bigint not null" fake:"skip"`
	Creator                        *User     `json:"creator" gorm:"foreignKey:CreatorId" fake:"skip"`
	MediaUploadTime                time.Time `json:"media_upload_time" gorm:"type:timestamp without time zone default current_timestamp" fake:"skip"`
}

func (nft *Nft) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	nft.Id = uuid
	return nil
}
