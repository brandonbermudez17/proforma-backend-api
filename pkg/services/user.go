package services

import (
	"proforma-backend-api/pkg/dto"
	"proforma-backend-api/pkg/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	DB *gorm.DB
}

type IUserService interface {
	Create(dto.UserInput) error
	Delete(uint) error
	Update(dto.UserInput, uint) error
	FetchAll(int, int) (dto.UserListOutput, error)
	GetByID(uint) (models.User, error)
}

func NewUserService(db *gorm.DB) IUserService {
	return &userService{
		DB: db,
	}
}

func (s *userService) Create(user dto.UserInput) error {
	person := models.PersonalData{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		DNI:       user.DNI,
		Email:     user.Email,
		Address:   user.Address,
	}

	pwd, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	return s.DB.Create(&models.User{
		AvatarPath:   &user.AvatarPath,
		Username:     user.Username,
		Password:     pwd,
		PersonalData: person,
		CompanyID:    user.CompanyID,
	}).Error
}
func (s *userService) Delete(id uint) error {
	return s.DB.Delete(models.User{}, id).Error
}

func (s *userService) Update(userData dto.UserInput, id uint) error {
	var user models.User
	if err := s.DB.First(&user, id).Error; err != nil {
		return err
	}
	user.AvatarPath = &userData.AvatarPath
	user.CompanyID = userData.CompanyID

	if user.Username != userData.Username {
		user.Username = userData.Username
	}

	if len(userData.Password) > 0 && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password)) != nil {
		var err error
		user.Username, err = HashPassword(userData.Password)
		if err != nil {
			return err
		}
	}

	user.PersonalData.FirstName = userData.FirstName
	user.PersonalData.LastName = userData.LastName
	user.PersonalData.Email = userData.Email
	user.PersonalData.Address = userData.Address
	user.PersonalData.DNI = userData.DNI

	return s.DB.Save(&user).Error

}
func (s *userService) FetchAll(limit int, offset int) (dto.UserListOutput, error) {
	var users []models.User
	var total int64

	if err := s.DB.Model(&models.Company{}).Count(&total).Error; err != nil {
		return dto.UserListOutput{}, err
	}

	err := s.DB.Limit(limit).Offset(offset).Find(&users).Error
	return dto.UserListOutput{
		Users: users,
		Total: uint(total),
	}, err
}

func (s *userService) GetByID(id uint) (models.User, error) {
	var user models.User
	err := s.DB.First(&user, id).Error
	return user, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}
