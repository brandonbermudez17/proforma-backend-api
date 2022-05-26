package services

import (
	"proforma-backend-api/pkg/dto"
	"proforma-backend-api/pkg/models"

	"gorm.io/gorm"
)

type clientService struct {
	DB *gorm.DB
}

type IClientService interface {
	Create(dto.ClientInput) error
	Delete(uint) error
	FetchAll(int, int) (dto.ClientListOutput, error)
	GetByID(uint) (models.Client, error)
	Update(dto.ClientInput, uint) error
}

func NewClientService(db *gorm.DB) IClientService {
	return &clientService{
		DB: db,
	}
}

func (s *clientService) Create(clientData dto.ClientInput) error {
	client := models.Client{
		Code:       clientData.Code,
		Name:       clientData.Name,
		Phone1:     clientData.Phone1,
		Phone2:     clientData.Phone2,
		Email:      clientData.Email,
		DNI:        clientData.DNI,
		Address:    clientData.Address,
		Country:    clientData.Country,
		PostalCode: clientData.PostalCode,
	}
	return s.DB.Create(&client).Error
}

func (s *clientService) Delete(id uint) error {
	return s.DB.Delete(&models.Client{}, id).Error
}

func (s *clientService) FetchAll(limit int, offset int) (dto.ClientListOutput, error) {
	var clients []models.Client
	var total int64

	if err := s.DB.Model(&models.Client{}).Count(&total).Error; err != nil {
		return dto.ClientListOutput{}, err
	}
	err := s.DB.Limit(limit).Offset(offset).Find(&clients).Error
	return dto.ClientListOutput{
		Clients: clients,
		Total:   uint(total),
	}, err
}
func (s *clientService) GetByID(id uint) (models.Client, error) {
	var client models.Client
	err := s.DB.First(&client, id).Error
	return client, err
}
func (s *clientService) Update(clientData dto.ClientInput, id uint) error {
	var client models.Client
	err := s.DB.First(&client, id).Error
	if err != nil {
		return err
	}

	client.Code = clientData.Code
	client.Name = clientData.Name
	client.Phone1 = clientData.Phone1
	client.Phone2 = clientData.Phone2
	client.Email = clientData.Email
	client.DNI = clientData.DNI
	client.Address = clientData.Address
	client.Country = clientData.Country
	client.PostalCode = clientData.PostalCode

	return s.DB.Save(&client).Error
}
