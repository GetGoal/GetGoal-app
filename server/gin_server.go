package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/config"
	"gorm.io/gorm"
)

type GinServer struct {
	app *gin.Engine
	db  *gorm.DB
	cfg *config.Config
}

// Start implements Server.
func (s *GinServer) Start() {

	serverURL := fmt.Sprintf(":%d", s.cfg.App.Port)

	s.app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s.app.Run(serverURL)
}

func NewGinServer(cfg *config.Config, db *gorm.DB) Server {
	return &GinServer{
		app: gin.New(),
		db:  db,
		cfg: cfg,
	}
}
