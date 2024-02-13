package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xbklyn/getgoal-app/exception"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
