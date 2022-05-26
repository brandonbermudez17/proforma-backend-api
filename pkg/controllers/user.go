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

type userController struct {
	Log         *zap.Logger
	UserService services.IUserService
}

func NewUserController(db *gorm.DB, log *zap.Logger) IUserController {
	return &userController{
		Log:         log,
		UserService: services.NewUserService(db),
	}
}

type IUserController interface {
	CreateOne() gin.HandlerFunc
	DeleteOne() gin.HandlerFunc
	Update() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	GetOne() gin.HandlerFunc
}

func (c *userController) CreateOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userData dto.UserInput
		if err := ctx.ShouldBindJSON(&userData); err != nil {
			c.Log.Error("Failed to read json", zap.Error(err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong parameters"})
			return
		}

		if err := c.UserService.Create(userData); err != nil {
			errMsg := "creation user failed"
			c.Log.Error(errMsg, zap.Error(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			return
		}

		ctx.IndentedJSON(http.StatusCreated, userData)
		c.Log.Info("created user successfully")
	}
}
func (c *userController) DeleteOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			errMsg := "invalid param"
			c.Log.Error(errMsg, zap.Error(err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			return
		}
		err = c.UserService.Delete(uint(id))
		if err != nil {
			errMsg := "record not found"
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			c.Log.Error(errMsg, zap.Error(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
func (c *userController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userData dto.UserInput

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			errMsg := "invalid param"
			c.Log.Error(errMsg, zap.Error(err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			return
		}

		if err := ctx.ShouldBindJSON(&userData); err != nil {
			c.Log.Error("failed to read json", zap.Error(err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.UserService.Update(userData, uint(id)); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
				return
			}

			errMsg := "update user failed"
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			c.Log.Error(errMsg, zap.Error(err))
			return
		}

		ctx.IndentedJSON(http.StatusOK, userData)
		c.Log.Info("user updated")
	}
}
func (c *userController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limit := ctx.GetInt("limit")
		offset := ctx.GetInt("offset")

		out, err := c.UserService.FetchAll(limit, offset)
		if err != nil {
			errMsg := "failed to get user records"
			c.Log.Error(errMsg, zap.Error(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			return
		}

		ctx.IndentedJSON(http.StatusOK, out)
	}
}
func (c *userController) GetOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			errMsg := "invalid param"
			c.Log.Error(errMsg, zap.Error(err))
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			return
		}
		data, err := c.UserService.GetByID(uint(id))
		if err != nil {
			errMsg := "failed to get user"
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			c.Log.Error(errMsg, zap.Error(err))
			return
		}

		ctx.IndentedJSON(http.StatusOK, data)
	}
}
