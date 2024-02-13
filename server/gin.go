package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/config"
	"github.com/xbklyn/getgoal-app/controller"
	"github.com/xbklyn/getgoal-app/model"
	repo "github.com/xbklyn/getgoal-app/repository/impl"
	service "github.com/xbklyn/getgoal-app/service/impl"
	"gorm.io/gorm"
)

type Gin struct {
	app *gin.Engine
	cfg *config.Config
	db  *gorm.DB
}

// Start implements Server.
func (s *Gin) Start() {

	//repo
	labelRepo := repo.NewlabelRepoImpl(s.db)

	//service
	labelService := service.NewLabelServiceImpl(&labelRepo)

	//controller
	labelController := controller.NewLabelController(labelService)

	serverURL := fmt.Sprintf(":%d", s.cfg.App.Port)

	//heatlh check
	s.app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.GeneralResponse{
			Code:    http.StatusOK,
			Message: "pong",
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

	//routes
	labelController.Route(v1)

	s.app.Run(serverURL)
}

func NewGinServer(cfg *config.Config, db *gorm.DB) Server {
	if cfg.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	return &Gin{
		app: gin.New(),
		db:  db,
		cfg: cfg,
	}
}