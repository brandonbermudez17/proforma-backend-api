package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Log    *zap.Logger
	Port   uint
	Host   string
	Engine *gin.Engine
}

func New(host string, port uint, db *gorm.DB, log *zap.Logger) Server {
	srv := Server{
		Engine: gin.New(),
		Host:   host,
		Port:   port,
		DB:     db,
		Log:    log,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	httpAddress := fmt.Sprintf("%s:%d", s.Host, s.Port)
	s.Log.Info(
		"Server running on",
		zap.String("host", s.Host),
		zap.Uint("port", s.Port),
	)

	return s.Engine.Run(httpAddress)
}
