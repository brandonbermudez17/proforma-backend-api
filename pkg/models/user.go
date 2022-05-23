package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string       `json:"username"`
	Password       string       `json:"-"`
	AvatarPath     *string      `json:"avatarPath"`
	DateOfLastSeen string       `json:"dateOfLastSeen"`
	Active         bool         `json:"active"`
	CompanyID      uint         `json:"-"`
	Company        Company      `json:"company"`
	PersonalDataID uint         `json:"-"`
	PersonalData   PersonalData `json:"personalData"`
}
