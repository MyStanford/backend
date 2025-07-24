package chat

import "github.com/gofiber/fiber/v2"

func EmotionGetRoute(ctx *fiber.Ctx) error {
	chatID := ctx.Query("chatID", "")
	if chatID == "" {
		return ctx.JSON(fiber.Map{
			"code":    1,
			"message": "missing chatID",
		})
	}
	nowChat, ok := AllChatDatas[chatID]
	if !ok {
		return ctx.JSON(fiber.Map{
			"code":    1,
			"message": "undefined chatID",
		})
	}
	data := make(map[string]int)
	for _, item := range nowChat.Persons {
		data[item.Name] = item.Emotion
	}
	return ctx.JSON(fiber.Map{
		"code":    0,
		"message": "",
		"data":    data,
	})
}
