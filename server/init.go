package server

import (
	"github.com/gofiber/fiber/v2"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitServer() {
	Server = fiber.New()
	Server.Use(fiberlogger.New(), recover.New())
}
