package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	LogoPath    string   `json:"logoPath"`
	DNI         string   `json:"dni"`
	Active      bool     `json:"active"`
	Clients     []Client `gorm:"many2many:client_company"`
	// CreatedByUser        uint
	// CreatedBy            User `gorm:"foreignKey:CreatedByUser"`
	// ModifiedByUser       uint
	// ModifiedBy           User `gorm:"foreignKey:ModifiedByUser"`
}
