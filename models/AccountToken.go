package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TokenType string

const (
	RESET_PASSWORD TokenType = "RESET_PASSWORD"
	VERIFY_EMAIL   TokenType = "VERIFY_EMAIL"
)

type AccountToken struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid primary key" fake:"skip"`
	TokenType TokenType `json:"media_type" sql:"type:MediaType" fake:"skip"`
	Token     string    `json:"token" gorm:"varchar(255) not null" fake:"skip"`
	UserId    uint64    `json:"user_id" gorm:"type:bigint" fake:"skip"`
	User      string    `json:"user" gorm:"foreignKey:UserId" fake:"skip"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp without time zone default current_timestamp" fake:"skip"`
	ExpireAt  time.Time `json:"expires_at" gorm:"type:timestamp without time zone default current_timestamp" fake:"skip"`
}

func (accountToken *AccountToken) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	accountToken.Id = uuid
	return nil
}
