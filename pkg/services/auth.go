package services

import (
	"proforma-backend-api/pkg/dto"
	"proforma-backend-api/pkg/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	DB *gorm.DB
}

type IAuthService interface {
	Login(dto.LoginDTO) error
}

func NewAuthService(db *gorm.DB) IAuthService {
	return &authService{
		DB: db,
	}
}

func (s *authService) Login(userData dto.LoginDTO) error {
	var user models.User
	err := s.DB.Where("username = ?", userData.Username).First(&user).Error
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))
}
