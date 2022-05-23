package server

import (
	"proforma-backend-api/pkg/controllers"
	"proforma-backend-api/pkg/middlewares"
)

func (s *Server) registerRoutes() {
	s.Engine.Use(middlewares.GinRecovery(s.Log, true), middlewares.GinLogger(s.Log))

	//initialize controllers
	authController := controllers.NewAuthController(s.DB, s.Log)

	//initialize group paths
	auth := s.Engine.Group("/auth")
	// private := s.Engine.Group("/api")

	auth.POST("/login", authController.Login())
}
