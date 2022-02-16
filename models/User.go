package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type RegisterChannel string

const (
	GOOGLE_LOGIN   RegisterChannel = "GOOGLE_LOGIN"
	EMAIL_REGISTER RegisterChannel = "EMAIL_REGISTER"
)

type UserType string

const (
	USER_TYPE_NORMAL  UserType = "USER_TYPE_NORMAL"
	USER_TYPE_CREATOR UserType = "USER_TYPE_CREATOR"
)

type UserStatus string

const (
	USER_STATUS_ACTIVE UserType = "USER_STATUS_ACTIVE"
	USER_STATUS_LOCKED UserType = "USER_STATUS_LOCKED"
)

type User struct {
	Id                   uint64          `json:"id" gorm:"type:bigserial primary key not null" fake:"skip"`
	RegisterChannel      RegisterChannel `json:"register_channel" sql:"type:RegisterChannel" fake:"skip"`
	UserType             UserType        `json:"user_type" gorm:"default:USER_TYPE_NORMAL" fake:"skip"`
	UserStatus           UserStatus      `json:"user_status" gorm:"default:USER_STATUS_ACTIVE" fake:"skip"`
	Verified             bool            `json:"verified" gorm:"type:bool default false" fake:"skip"`
	FirstName            string          `json:"first_name" gorm:"type:varchar(100)" fake:"{firstname}"`
	LastName             string          `json:"last_name" gorm:"type:varchar(100)" fake:"{lastname}"`
	DisplayName          string          `json:"display_name" gorm:"type:varchar(250)" fake:"{username}"`
	UserName             string          `json:"user_name" gorm:"type:varchar(250) unique" fake:"{username}"`
	HashedPassword       string          `json:"hashed_password" gorm:"type:varchar(255)" fake:"{password}"`
	PhoneNumber          string          `json:"phone_number" gorm:"type:varchar(250)" fake:"{phone}"`
	EmailAddress         string          `json:"email_address" gorm:"type:varchar(255) not null unique" fake:"{email}"`
	AvatarUrl            string          `json:"avatar_url" gorm:"type:varchar(100) default ''" fake:"{url}"`
	Bio                  string          `json:"bio" gorm:"type:varchar(1000) default ''" fake:"{sentence:5}"`
	YoutubeUrl           string          `json:"youtube_url" gorm:"type:varchar(255)" fake:"{url}"`
	TiktokUrl            string          `json:"tiktok_url" gorm:"type:varchar(255)" fake:"{url}"`
	InstagramUrl         string          `json:"instagram_url" gorm:"type:varchar(255)" fake:"{url}"`
	TwitchUrl            string          `json:"twitch_url" gorm:"type:varchar(255)" fake:"{url}"`
	TwitterUrl           string          `json:"twitter_url" gorm:"type:varchar(255)" fake:"{url}"`
	IntroYoutubeEmbedUrl string          `json:"intro_youtube_embed_url" gorm:"type:varchar(255)" fake:"{url}"`
	LastLoginAt          time.Time       `json:"last_login_at" sql:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP" fake:"skip"`
	UpdatedAt            time.Time       `json:"updated_at" gorm:"type:timestamp" fake:"skip"`
	CreatedAt            time.Time       `json:"created_at" gorm:"type:timestamp without time zone" fake:"skip"`
	Tags                 pq.StringArray  `json:"tags" gorm:"type:varchar(255)[]" fakesize:"2"`
	InternalTags         pq.StringArray  `json:"internal_tags" gorm:"type:varchar(255)[]" fakesize:"2"`
	ChainWalletId        uuid.UUID       `json:"chain_wallet_id" gorm:"type:uuid" fake:"skip"`
	ChainWallet          *ChainWallet    `json:"chain_wallet" gorm:"foreignKey:ChainWalletId" fake:"skip"`
}
