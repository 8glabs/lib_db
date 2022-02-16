package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ChainWallet struct {
	Id                             uuid.UUID      `json:"id" gorm:"type:uuid primary key"`
	OwnerId                        uint64         `json:"owner_id" gorm:"type:bigint"`
	CustodialStarkexPublicKey      string         `json:"custodial_starkex_public_key"`
	CustodialStarkexPrivateKey     string         `json:"custodial_starkex_private_key"`
	CustodialStarkexNextVaultId    uint64         `json:"custodial_starkex_next_vault_id"`
	CustodialEthereumWalletAddress string         `json:"custodial_ethereum_wallet_address"`
	CustodialEthereumPublicKey     string         `json:"custodial_ethereum_public_key"`
	CustodialEthereumPrivateKey    string         `json:"custodial_ethereum_private_key"`
	CustodialEthereumMnemonics     pq.StringArray `gorm:"type:integer[]"`
	CustodialSolanaPrivateKeyStr   string         `json:"custodial_solana_privatekey_str" gorm:"type:varchar(1000)"`
}

func (chainWallets *ChainWallet) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	chainWallets.Id = uuid
	return nil
}
