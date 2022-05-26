package controllers

import (
	"errors"
	"net/http"
	"proforma-backend-api/pkg/dto"
	"proforma-backend-api/pkg/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type companyController struct {
	Log            *zap.Logger
	CompanyService services.ICompanyService
}

func NewCompanyController(db *gorm.DB, log *zap.Logger) ICompanyController {
	return &companyController{
		Log:            log,
		CompanyService: services.NewCompanyService(db),
	}
}

type ICompanyController interface {
	CreateOne() gin.HandlerFunc
	DeleteOne() gin.HandlerFunc
	Update() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	GetOne() gin.HandlerFunc
}

func (c *companyController) CreateOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var companyData dto.CompanyInput
		if err := ctx.ShouldBindJSON(&companyData); err != nil {
			c.Log.Error("Failed to read json", zap.Error(err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.CompanyService.Create(companyData); err != nil {
			errMsg := "failed to add company"
			c.Log.Error(errMsg, zap.Error(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			return
		}

		ctx.IndentedJSON(http.StatusCreated, companyData)
	}
}
func (c *companyController) DeleteOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			errMsg := "invalid param"
			c.Log.Error(errMsg, zap.Error(err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			return
		}
		err = c.CompanyService.Delete(uint(id))
		if err != nil {
			errMsg := "record not found"
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			c.Log.Error(errMsg, zap.Error(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func (c *companyController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var companyData dto.CompanyInput

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			errMsg := "invalid param"
			c.Log.Error(errMsg, zap.Error(err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			return
		}

		if err := ctx.ShouldBindJSON(&companyData); err != nil {
			c.Log.Error("failed to read json", zap.Error(err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.CompanyService.Update(companyData, uint(id)); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
				return
			}

			errMsg := "Failed to update company data"
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			c.Log.Error(errMsg, zap.Error(err))
			return
		}

		ctx.IndentedJSON(http.StatusOK, companyData)
	}
}
func (c *companyController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limit := ctx.GetInt("limit")
		offset := ctx.GetInt("offset")

		out, err := c.CompanyService.FetchAll(limit, offset)
		if err != nil {
			errMsg := "failed to get company records"
			c.Log.Error(errMsg, zap.Error(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			return
		}

		ctx.IndentedJSON(http.StatusOK, out)
	}
}
func (c *companyController) GetOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			errMsg := "invalid param"
			c.Log.Error(errMsg, zap.Error(err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			return
		}
		data, err := c.CompanyService.GetByID(uint(id))
		if err != nil {
			errMsg := "failed to get company"
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			c.Log.Error(errMsg, zap.Error(err))
			return
		}

		ctx.IndentedJSON(http.StatusOK, data)
	}
}
