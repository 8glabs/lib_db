package models

import "time"

type PaymentHistory struct {
	Id            uint64    `json:"id" gorm:"type:bigserial primary key not null" fake:"skip"`
	PayerId       uint64    `json:"payer_id"`
	PaymentTool   string    `json:"payment_tool" gorm:"type:varchar(20)"`
	StripeOrderId string    `json:"stripe_order_id" gorm:"type:varchar(255)"`
	Amount        float64   `json:"amount" gorm:"type:double precision"`
	Currency      string    `json:"currency" gorm:"type:varchar(20) default 'USD'"`
	Payer         *User     `json:"payer" gorm:"foreignKey:PayerId" fake:"skip"`
	CreatedAt     time.Time `json:"created_at" gorm:"timestamp without time zone default current_timestamp"`
}
