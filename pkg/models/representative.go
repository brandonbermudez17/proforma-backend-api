package models

import "gorm.io/gorm"

type Representative struct {
	gorm.Model
	PersonalDataID uint         `json:"-"`
	PersonalData   PersonalData `json:"personalData"`
}
