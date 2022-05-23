package dto

import "proforma-backend-api/pkg/models"

type ClientInput struct {
	Code       string `json:"code" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Phone1     string `json:"phone1" binding:"required"`
	Phone2     string `json:"phone2" binding:"required"`
	Email      string `json:"email" binding:"required"`
	DNI        string `json:"dni" binding:"required"`
	Address    string `json:"address" binding:"required"`
	Country    string `json:"country" binding:"required"`
	PostalCode string `json:"postalCode" binding:"required"`
}

type ClientListOutput struct {
	Clients []models.Client `json:"clients"`
	Total   uint            `json:"total"`
}
