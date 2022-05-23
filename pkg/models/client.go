package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Code       string
	Name       string
	Phone1     string
	Phone2     string
	Email      string
	DNI        string
	Address    string
	Country    string
	PostalCode string
	Active     bool
	Companies  []Company `gorm:"many2many:client_company"`
}
