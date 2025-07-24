package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	HttpDefaultMap = fiber.Map{
		"code":    0,
		"message": "",
	}
)

func HttpBodyParseCheck(ctx *fiber.Ctx, data interface{}) error {
	e := ctx.BodyParser(data)
	if e != nil {
		ctx.JSON(fiber.Map{
			"code":    1,
			"message": e.Error(),
		})
		return e
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(data)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			ctx.JSON(fiber.Map{
				"code":    1,
				"message": err.Error(),
			})
			return err
		}

		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {
				ctx.JSON(fiber.Map{
					"code":    1,
					"message": e.Field() + " error",
				})
				return e
			}
		}
	}
	return nil
}
