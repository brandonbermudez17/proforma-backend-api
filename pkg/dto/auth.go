package dto

import "proforma-backend-api/pkg/models"

type LoginDTO struct {
	Username string
	Password string
}

type UserInput struct {
	FirstName  string   `json:"firstName"`
	LastName   string   `json:"lastName"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	DNI        string   `json:"dni"`
	Email      string   `json:"email"`
	Address    string   `json:"address"`
	AvatarPath string   `json:"avatarPath"`
	CompanyID  uint     `json:"companyId"`
	Phones     []string `json:"phones"`
}

type UserListOutput struct {
	Users []models.User `json:"users"`
	Total uint          `json:"total"`
}
