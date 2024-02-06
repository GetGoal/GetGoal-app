package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xbklyn/getgoal-app/config"
	ahandlers "github.com/xbklyn/getgoal-app/handlers/action"
	lhandlers "github.com/xbklyn/getgoal-app/handlers/label"
	arepositories "github.com/xbklyn/getgoal-app/repositories/action"
	lrepositories "github.com/xbklyn/getgoal-app/repositories/label"
	ausecases "github.com/xbklyn/getgoal-app/usecases/action"
	lusecases "github.com/xbklyn/getgoal-app/usecases/label"
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
	s.initializeActionHandler(v1)
	s.app.Run(serverURL)
}

func NewGinServer(cfg *config.Config, db *gorm.DB) Server {
	if cfg.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	return &ginServer{
		app: gin.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *ginServer) initializeLabelHandler(v1 *gin.RouterGroup) {
	//Init all layers
	repo := lrepositories.NewLabelRepositoryImpl(s.db)
	usecase := lusecases.NewLabelUsecaseImpl(repo)
	handler := lhandlers.NewLabelHandlerImpl(usecase)

	labelRouter := v1.Group("/labels")
	labelRouter.GET("", handler.FindAllLabels)
	labelRouter.GET("/:id", handler.FindLabelByID)
	labelRouter.GET("/search", handler.GetSeachLabel)
	labelRouter.POST("", handler.AddNewLabel)
}

func (s *ginServer) initializeActionHandler(v1 *gin.RouterGroup) {
	//Init all layers
	repo := arepositories.NewActionRepositoryImpl(s.db)
	usecase := ausecases.NewActionUsecaseImpl(repo)
	handler := ahandlers.NewActionHandlerImpl(usecase)

	actionRouter := v1.Group("/actions")
	actionRouter.GET("", handler.FindAllActions)
	actionRouter.GET("/:id", handler.FindActionByID)
}
