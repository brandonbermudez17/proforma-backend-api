package server

import (
	"proforma-backend-api/pkg/controllers"
	"proforma-backend-api/pkg/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) registerRoutes() {
	s.Engine.Use(
		middlewares.GinRecovery(s.Log, true),
		middlewares.GinLogger(s.Log),
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "GET"},
			AllowCredentials: true,
		}))

	//initialize controllers
	authController := controllers.NewAuthController(s.DB, s.Log)
	userController := controllers.NewUserController(s.DB, s.Log)
	companyController := controllers.NewCompanyController(s.DB, s.Log)
	clientController := controllers.NewClientController(s.DB, s.Log)

	//initialize group paths
	auth := s.Engine.Group("/auth")
	api := s.Engine.Group("/api")

	auth.POST("/login", authController.Login())

	createUserEndpoints(api, userController)
	createClientEndpoints(api, clientController)
	createCompanyEndpoints(api, companyController)
}

func createUserEndpoints(api *gin.RouterGroup, c controllers.IUserController) {
	userEnpoints := api.Group("/user")
	userEnpoints.GET("/:id", c.GetOne())
	userEnpoints.DELETE("/:id", c.DeleteOne())
	userEnpoints.PUT("/:id", c.Update())
	userEnpoints.POST("/", c.CreateOne())
	userEnpoints.GET("/list", c.GetAll())
}

func createClientEndpoints(api *gin.RouterGroup, c controllers.IClientController) {
	clientEnpoints := api.Group("/client")
	clientEnpoints.GET("/:id", c.GetOne())
	clientEnpoints.DELETE("/:id", c.DeleteOne())
	clientEnpoints.PUT("/:id", c.Update())
	clientEnpoints.POST("/", c.CreateOne())
	clientEnpoints.GET("/list", c.GetAll())
}

func createCompanyEndpoints(api *gin.RouterGroup, c controllers.ICompanyController) {
	companyEnpoints := api.Group("/company")
	companyEnpoints.GET("/:id", c.GetOne())
	companyEnpoints.DELETE("/:id", c.DeleteOne())
	companyEnpoints.PUT("/:id", c.Update())
	companyEnpoints.POST("/", c.CreateOne())
	companyEnpoints.GET("/list", c.GetAll())
}
