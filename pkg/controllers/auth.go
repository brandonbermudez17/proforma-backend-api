package controllers

import (
	"net/http"
	"proforma-backend-api/pkg/dto"
	"proforma-backend-api/pkg/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type authController struct {
	Log         *zap.Logger
	AuthService services.IAuthService
}

func NewAuthController(db *gorm.DB, log *zap.Logger) IAuthController {
	return &authController{
		Log:         log,
		AuthService: services.NewAuthService(db),
	}
}

type IAuthController interface {
	Login() gin.HandlerFunc
}

func (c *authController) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userData dto.LoginDTO
		if err := ctx.BindJSON(&userData); err != nil {
			c.Log.Error("failed to read json", zap.Error(err))
			return
		}

		if err := c.AuthService.Login(userData); err != nil {
			errMsg := "failed to login user"
			c.Log.Error(errMsg, zap.Error(err))
			ctx.JSON(http.StatusNotFound, gin.H{"error": errMsg})
			return
		}

		msgSuccess := "Successfully authenticated user"
		ctx.JSON(http.StatusOK, gin.H{"message": msgSuccess})
		c.Log.Info(msgSuccess)
	}
}
