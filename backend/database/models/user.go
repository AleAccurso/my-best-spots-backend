package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID  `json:"id" gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;default:now();not null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;default:now();not null"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`

	Nickname    string    `json:"nickname" gorm:"column:nickname;type:varchar(64);not null"`
	Gender      string    `json:"gender" gorm:"column:gender;type:varchar(6);not null"`
	Email       string    `json:"email" gorm:"column:email;type:varchar(128);unique_index;not null"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number;type:varchar(64);not null"`
	AddressId   uuid.UUID `json:"address_id" gorm:"column:address_id"`
	Address     Address
	Role        string     `json:"role" gorm:"column:role;type:varchar(5);not null"`
	ExternalId  string     `json:"external_id" gorm:"column:external_id;type:varchar(64);not null"`
	VerifiedAt  *time.Time `json:"verified_at" gorm:"column:verified_at"`
	LastLoginAt *time.Time `json:"last_login_at" gorm:"column:last_login_at"`
}
