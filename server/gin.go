package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	file "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xbklyn/getgoal-app/config"
	"github.com/xbklyn/getgoal-app/controller"
	_ "github.com/xbklyn/getgoal-app/docs"
	"github.com/xbklyn/getgoal-app/middleware"
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
	taskRepo := repo.NewTaskRepoImpl(s.db)
	programRepo := repo.NewProgramRepoImpl(s.db)
	userRepo := repo.NewUserRepoImpl(s.db)
	userProgramRepo := repo.NewUserProgramRepoImpl(s.db)

	//service
	labelService := service.NewLabelServiceImpl(&labelRepo)
	taskService := service.NewTaskServiceImpl(taskRepo, userRepo, userProgramRepo)
	programService := service.NewProgramServiceImpl(programRepo, taskRepo, labelRepo, userRepo, userProgramRepo)
	mailerService := service.NewMailerServiceImpl()
	authService := service.NewAuthServiceImpl(userRepo, mailerService)
	//controller

	labelController := controller.NewLabelController(labelService)
	taskController := controller.NewTaskController(taskService)
	programController := controller.NewProgramController(programService)
	authController := controller.NewAuthController(authService)

	serverURL := fmt.Sprintf(":%d", s.cfg.App.Port)

	env := s.cfg.Env
	s.app.GET(env+"/swagger/*any", ginSwagger.WrapHandler(file.Handler))

	// HealthCheckHandler godoc
	// @summary Health Check
	// @description Health checking for the service
	// @id HealthCheckHandler
	// @produce json
	// @response 200 {string} string "OK"
	// @router /healthcheck [get]
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

	//No header required
	authController.RouteAnonymous(v1)

	//Enable middleware
	v1.Use(middleware.JWTAuthMiddleware(authService.(*service.AuthServiceImpl), []byte(s.cfg.JwtKeys.AccessSecret)))

	//Header required
	labelController.Route(v1)
	taskController.Route(v1)
	programController.Route(v1)
	authController.Route(v1)

	s.app.Run(serverURL)
}

func NewGinServer(cfg *config.Config, db *gorm.DB) Server {
	if cfg.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	return &Gin{
		app: gin.Default(),
		db:  db,
		cfg: cfg,
	}
}
