package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PrimarySaleCampaign struct {
	Id                   uuid.UUID  `json:"id" gorm:"type:uuid primary key" fake:"skip"`
	PrimarySaleStartTime *time.Time `json:"primary_sale_start_time" gorm:"type:timestamp without time zone default current_timestamp" fake:"skip"`
	PrimarySaleEndTime   *time.Time `json:"primary_sale_end_time" gorm:"type:timestamp without time zone default current_timestamp" fake:"skip"`
	Campaign
}

func (campaign *PrimarySaleCampaign) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	campaign.Id = uuid
	return nil
}
