package services

import (
	"proforma-backend-api/pkg/dto"
	"proforma-backend-api/pkg/models"

	"gorm.io/gorm"
)

type companyService struct {
	DB *gorm.DB
}

type ICompanyService interface {
	Create(dto.CompanyInput) error
	Delete(uint) error
	FetchAll(int, int) (dto.CompanyListOutput, error)
	GetByID(uint) (models.Company, error)
	Update(dto.CompanyInput, uint) error
}

func NewCompanyService(db *gorm.DB) ICompanyService {
	return &companyService{
		DB: db,
	}
}

func (s *companyService) Create(companyData dto.CompanyInput) error {
	company := models.Company{
		Name:        companyData.Name,
		Description: companyData.Description,
		LogoPath:    companyData.LogoPath,
		DNI:         companyData.DNI,
	}
	return s.DB.Create(&company).Error
}

func (s *companyService) Delete(id uint) error {
	return s.DB.Delete(&models.Company{}, id).Error
}

func (s *companyService) FetchAll(limit int, offset int) (dto.CompanyListOutput, error) {
	var companies []models.Company
	var total int64

	if err := s.DB.Model(&models.Company{}).Count(&total).Error; err != nil {
		return dto.CompanyListOutput{}, err
	}

	err := s.DB.Limit(limit).Offset(offset).Find(&companies).Error
	return dto.CompanyListOutput{
		Companies: companies,
		Total:     uint(total),
	}, err
}

func (s *companyService) GetByID(id uint) (models.Company, error) {
	var company models.Company
	err := s.DB.First(&company, id).Error
	return company, err
}

func (s *companyService) Update(companyData dto.CompanyInput, id uint) error {
	var company models.Company
	if err := s.DB.First(&company, id).Error; err != nil {
		return err
	}
	company.Name = companyData.Name
	company.Description = companyData.Description
	company.LogoPath = companyData.LogoPath
	company.DNI = companyData.DNI

	return s.DB.Save(&company).Error
}
