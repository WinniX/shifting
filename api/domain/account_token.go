package domain

import (
	"time"

	"gorm.io/gorm"
)

type AccountToken struct {
	gorm.Model
	AccessToken    string    `gorm:"not null"`
	RefreshToken   string    `gorm:"not null"`
	Expiry         time.Time `gorm:"type:timestamptz;not null"`
	AccountCode    string    `gorm:"not null;size:10;uniqueIndex"`
	ApaleoOneToken string    `gorm:"not null;uniqueIndex"`
}
