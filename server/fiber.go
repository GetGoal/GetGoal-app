package server

// import (
// 	"encoding/json"
// 	"fmt"
// 	"strconv"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/xbklyn/getgoal-app/config"
// 	"github.com/xbklyn/getgoal-app/controller"
// 	"github.com/xbklyn/getgoal-app/exception"
// 	repo "github.com/xbklyn/getgoal-app/repository/impl"
// 	service "github.com/xbklyn/getgoal-app/service/impl"
// 	"gorm.io/gorm"
// )

// type Fiber struct {
// 	app *fiber.App
// 	cfg *config.Config
// 	db  *gorm.DB
// }

// // Start implements Server.
// func (s *Fiber) Start() {

// 	//repo
// 	labelRepo := repo.NewlabelRepoImpl(s.db)

// 	//service
// 	labelService := service.NewLabelServiceImpl(&labelRepo)

// 	//controller
// 	labelController := controller.NewLabelController(labelService)

// 	//heatlh check
// 	s.app.Get("/ping", func(c *fiber.Ctx) error {
// 		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "pong"})
// 	}).Name("ping")

// 	//api group
// 	api := s.app.Group("/api/" + func() string {
// 		env := s.cfg.Env
// 		if env == "prod" {
// 			return ""
// 		}
// 		return env
// 	}())

// 	v1 := api.Group("/v1")

// 	//routes
// 	labelController.Route(v1)

// 	routes := s.app.GetRoutes(true)

// 	for _, route := range routes {
// 		str, _ := json.MarshalIndent(route, "", "  ")

// 		fmt.Println(string(str))
// 	}
// 	// Prints:
// 	// {
// 	//    "method": "GET",
// 	//    "name": "api",
// 	//    "path": "/api/*",
// 	//    "params": [
// 	//      "*1"
// 	//    ]
// 	// }
// 	port := ":" + strconv.Itoa(s.cfg.App.Port)
// 	err := s.app.Listen(port)
// 	exception.PanicLogging(err)
// }

// func FiberServer(cfg *config.Config, db *gorm.DB) Server {

// 	return &Fiber{
// 		app: fiber.New(config.NewFiberConfiguration()),
// 		cfg: cfg,
// 		db:  db,
// 	}
// }
