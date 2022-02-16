package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GiftingCampaign struct {
	Id                    uuid.UUID  `json:"id" gorm:"type:uuid primary key" fake:"skip"`
	GiftingStartTime      *time.Time `json:"gifting_start_time" gorm:"type:timestamp without time zone default current_timestamp" fake:"skip"`
	GiftingEndTime        *time.Time `json:"gifting_end_time" gorm:"type:timestamp without time zone default current_timestamp" fake:"skip"`
	NextAvailableSerialId int        `json:"next_available_serial_id" gorm:"type:integer not null default 1" fake:"skip"`
	Campaign
}

func (campaign *GiftingCampaign) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	campaign.Id = uuid
	return nil
}
