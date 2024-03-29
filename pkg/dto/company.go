package dto

import "proforma-backend-api/pkg/models"

type CompanyInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	LogoPath    string `json:"logoPath" binding:"required"`
	DNI         string `json:"dni" binding:"required"`
}

type CompanyListOutput struct {
	Companies []models.Company `json:"companies"`
	Total     uint             `json:"total"`
}
