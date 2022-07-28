package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Shift struct {
	gorm.Model
	UID         uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	ValidFrom   time.Time    `gorm:"type:timestamp;not null"`
	ValidUntil  sql.NullTime `gorm:"type:timestamp"`
	WeekDay     time.Weekday `gorm:"not null"`
	StartsAt    time.Time    `gorm:"type:timestamp;not null"`
	EndsAt      time.Time    `gorm:"type:timestamp;not null"`
	RequiresPpl int          `gorm:"not null"`
	AccountCode string       `gorm:"not null;size:10"`
	PropertyId  string       `gorm:"not null;size:10"`
}

type Employee struct {
	gorm.Model
	UID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	AccountCode string    `gorm:"not null;size:10"`
	PropertyId  string    `gorm:"not null;size:10"`
}

type EmployeeShift struct {
	gorm.Model
	Employee   Employee
	EmployeeID uint `gorm:"not null"`
	Shift      Shift
	ShiftID    uint `gorm:"not null"`
}
