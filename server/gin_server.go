package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/config"
	"github.com/xbklyn/getgoal-app/handlers"
	"github.com/xbklyn/getgoal-app/repositories"
	"github.com/xbklyn/getgoal-app/usecases"
	"gorm.io/gorm"
)

type ginServer struct {
	app *gin.Engine
	db  *gorm.DB
	cfg *config.Config
}

// Start implements Server.
func (s *ginServer) Start() {

	serverURL := fmt.Sprintf(":%d", s.cfg.App.Port)

	//heatlh check
	s.app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//api group
	api := s.app.Group("/api/" + func() string {
		env := s.cfg.Env
		if env == "prod" {
			return ""
		}
		return env
	}())

	v1 := api.Group("/v1")

	s.initializeLabelHandler(v1)
	s.app.Run(serverURL)
}

func NewGinServer(cfg *config.Config, db *gorm.DB) Server {
	return &ginServer{
		app: gin.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *ginServer) initializeLabelHandler(v1 *gin.RouterGroup) {
	//Init all layers
	repo := repositories.NewLabelRepositoryImpl(s.db)
	usecase := usecases.NewLabelUsecaseImpl(repo)
	handler := handlers.NewLabelHandlerImpl(usecase)

	labelRouter := v1.Group("/labels")
	labelRouter.GET("", handler.FindAllLabels)
}
