package person

import (
	"mystanford/database"
	"mystanford/person"
	"mystanford/utils"

	"github.com/gofiber/fiber/v2"
)

type AddRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Prompt      string `json:"prompt" form:"prompt" validate:"required"`
}

func PersonAddRoute(ctx *fiber.Ctx) error {
	var bodyData AddRequest
	e := utils.HttpBodyParseCheck(ctx, &bodyData)
	if e != nil {
		return nil
	}
	if database.PersonExist(bodyData.Name) {
		return ctx.JSON(fiber.Map{
			"code":    2,
			"message": "人物已存在",
		})
	}
	e = database.PersonAdd(person.Person{
		Name:        bodyData.Name,
		Description: bodyData.Description,
		Prompt:      bodyData.Prompt,
	})
	if e != nil {
		return ctx.JSON(fiber.Map{
			"code":    2,
			"message": e.Error(),
		})
	}
	return ctx.JSON(utils.HttpDefaultMap)
}
