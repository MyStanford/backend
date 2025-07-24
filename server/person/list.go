package person

import (
	"mystanford/database"

	"github.com/gofiber/fiber/v2"
)

func PersonListRoute(ctx *fiber.Ctx) error {
	persons, e := database.PersonGetAll()
	if e != nil {
		return ctx.JSON(fiber.Map{
			"code":    2,
			"message": e.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"code":    0,
		"message": "",
		"data":    persons,
	})
}
