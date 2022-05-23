package models

import "gorm.io/gorm"

type PersonalData struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DNI       string `json:"dni"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	// CommentaryOfDeletion string  `json:"commentaryOfDeletion"`
	Phone1 string
	Phone2 string
}
