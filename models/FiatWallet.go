package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FiatWallet struct {
	Id              uuid.UUID `json:"id" gorm:"type:uuid primary key"`
	OwnerId         uint64    `json:"owner_id" gorm:"type:bigint not null"`
	BalanceCurrency string    `json:"balance_currency" gorm:"type:varchar(50)"`
	Balance         float64   `json:"balance" gorm:"type:double precision default 0"`
}

func (fiatWallet *FiatWallet) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	fiatWallet.Id = uuid
	return nil
}
