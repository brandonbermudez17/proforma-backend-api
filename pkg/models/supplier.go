package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Code       string
	Name       string
	TradeName  string
	Email      string
	Phone1     string
	Phone2     string
	Nif        string
	Address    string
	PostalCode string
	Country    string
}
