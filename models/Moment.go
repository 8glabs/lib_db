package models

type MediaType string

const (
	MEDIA_TYPE_VIDEO MediaType = "VIDEO"
	MEDIA_TYPE_IMAGE MediaType = "IMAGE"
	MEDIA_TYPE_GIF   MediaType = "GIF"
)

// type Moment struct {
// 	Id uuid.UUID `json:"id" gorm:"type:uuid primary key" fake:"skip"`

// 	AutographUrl string         `json:"autograph_url" gorm:"varchar(255)" fake:"{url}"`
// 	Tags         pq.StringArray `json:"tags" gorm:"type:varchar(255)[]" fakesize:"2"`
// 	InternalTags pq.StringArray `json:"internal_tags" gorm:"type:varchar(255)[]" fakesize:"2"`
// }

// func (moment *Moment) BeforeCreate(tx *gorm.DB) error {
// 	uuid, err := uuid.NewRandom()
// 	if err != nil {
// 		return err
// 	}
// 	moment.Id = uuid
// 	return nil
// }
