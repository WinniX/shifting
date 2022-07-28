package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRequest struct {
	gorm.Model
	State uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
}
