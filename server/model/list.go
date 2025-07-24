package model

import (
	"mystanford/config"

	"github.com/gofiber/fiber/v2"
)

func ModelListRoute(ctx *fiber.Ctx) error {
	data := []string{}
	for _, i := range config.NowConfig.Models {
		data = append(data, i.Name)
	}
	return ctx.JSON(fiber.Map{
		"code":    0,
		"message": "",
		"data":    data,
	})
}
