package models

import "time"

// UserCredential handles authentication data (Table 1)
type UserCredential struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex;not null;size:50"`
	Email     string `gorm:"uniqueIndex;not null;size:100"`
	Password  string `gorm:"not null"` // Hashed password
	CreatedAt time.Time
	UpdatedAt time.Time
	Profile   UserProfile `gorm:"foreignKey:CredentialID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// UserProfile handles the master data (Table 2)
type UserProfile struct {
	ID           uint   `gorm:"primaryKey"`
	CredentialID uint   `gorm:"uniqueIndex;not null"` // Foreign Key to UserCredential
	FullName     string `gorm:"size:100;not null"`
	PhoneNumber  string `gorm:"size:20"`
	Address      string `gorm:"type:text"`
}
