package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *sql.NullTime `gorm:"index"`
}
