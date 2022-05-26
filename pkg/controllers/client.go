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

type clientController struct {
	Log           *zap.Logger
	ClientService services.IClientService
}

func NewClientController(db *gorm.DB, log *zap.Logger) IClientController {
	return &clientController{
		Log:           log,
		ClientService: services.NewClientService(db),
	}
}

type IClientController interface {
	CreateOne() gin.HandlerFunc
	DeleteOne() gin.HandlerFunc
	Update() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	GetOne() gin.HandlerFunc
}

func (c *clientController) CreateOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var clientData dto.ClientInput
		if err := ctx.ShouldBindJSON(&clientData); err != nil {
			c.Log.Error("Failed to read json", zap.Error(err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.ClientService.Create(clientData); err != nil {
			errMsg := "failed to add client"
			c.Log.Error(errMsg, zap.Error(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			return
		}

		ctx.IndentedJSON(http.StatusCreated, clientData)
	}
}
func (c *clientController) DeleteOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			errMsg := "invalid param"
			c.Log.Error(errMsg, zap.Error(err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			return
		}
		err = c.ClientService.Delete(uint(id))
		if err != nil {
			errMsg := "record not found"
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			c.Log.Error(errMsg, zap.Error(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func (c *clientController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var clientData dto.ClientInput

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			errMsg := "invalid param"
			c.Log.Error(errMsg, zap.Error(err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			return
		}

		if err := ctx.ShouldBindJSON(&clientData); err != nil {
			c.Log.Error("failed to read json", zap.Error(err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.ClientService.Update(clientData, uint(id)); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
				return
			}

			errMsg := "Failed to update client data"
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			c.Log.Error(errMsg, zap.Error(err))
			return
		}

		ctx.IndentedJSON(http.StatusOK, clientData)
	}
}
func (c *clientController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limit := ctx.GetInt("limit")
		offset := ctx.GetInt("offset")

		out, err := c.ClientService.FetchAll(limit, offset)
		if err != nil {
			errMsg := "failed to get client records"
			c.Log.Error(errMsg, zap.Error(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			return
		}

		ctx.IndentedJSON(http.StatusOK, out)
	}
}
func (c *clientController) GetOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			errMsg := "invalid param"
			c.Log.Error(errMsg, zap.Error(err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			return
		}
		data, err := c.ClientService.GetByID(uint(id))
		if err != nil {
			errMsg := "failed to get client"
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			c.Log.Error(errMsg, zap.Error(err))
			return
		}

		ctx.IndentedJSON(http.StatusOK, data)
	}
}
