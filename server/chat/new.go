package chat

import (
	"context"
	"mystanford/config"
	"mystanford/database"
	"mystanford/person"
	"mystanford/utils"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/gofiber/fiber/v2"
)

type NewRequest struct {
	Model  string   `json:"model" form:"model" validate:"required"`
	People []string `json:"people" form:"people" validate:"required"`
}

func NewChatRoute(ctx *fiber.Ctx) error {
	var bodyData NewRequest
	e := utils.HttpBodyParseCheck(ctx, &bodyData)
	if e != nil {
		return nil
	}
	if len(bodyData.People) > 5 || len(bodyData.People) < 1 {
		return ctx.JSON(fiber.Map{
			"code":    2,
			"message": "人物数量不合法",
		})
	}
	newChat := Chat{
		Persons:         []*person.Person{},
		HistoryMessages: []*schema.Message{},
	}
	flag := false
	for _, item := range config.NowConfig.Models {
		if item.Name == bodyData.Model {
			newChat.Model = item
			flag = true
			break
		}
	}
	if !flag {
		return ctx.JSON(fiber.Map{
			"code":    1,
			"message": "model:" + bodyData.Model + "不存在",
		})
	}
	for _, item := range bodyData.People {
		p, e := database.PersonGetByName(item)
		if e != nil {
			return ctx.JSON(fiber.Map{
				"code":    1,
				"message": "person:" + item + "不存在",
			})
		}
		newChat.Persons = append(newChat.Persons, &person.Person{
			Name:        p.Name,
			Description: p.Description,
			Prompt:      p.Prompt,
			Emotion:     0,
		})
	}
	newChat.ChatID = utils.UUID()
	var agentModel model.ToolCallingChatModel
	switch newChat.Model.Type {
	case "openai":
		agentModel, _ = openai.NewChatModel(context.Background(), &openai.ChatModelConfig{
			BaseURL: newChat.Model.Path,
			APIKey:  newChat.Model.Key,
			Model:   newChat.Model.Model,
		})
	default:
		return ctx.JSON(fiber.Map{
			"code":    1,
			"message": "不支持的模型类型",
		})
	}
	newSystemPrompt := systemPrompt
	for _, item := range newChat.Persons {
		newSystemPrompt += "\n\n姓名：" + item.Name + "\n简介：" + item.Description + "\n具体信息：" + item.Prompt
	}
	newChat.Agent, e = react.NewAgent(context.Background(), &react.AgentConfig{
		ToolCallingModel: agentModel,
		ToolsConfig: compose.ToolsNodeConfig{Tools: []tool.BaseTool{
			&EmotionTool{
				ChatID: newChat.ChatID,
			},
		}},
	})
	newChat.HistoryMessages = append(newChat.HistoryMessages, &schema.Message{
		Role:    "system",
		Content: newSystemPrompt,
	})
	if e != nil {
		return ctx.JSON(fiber.Map{
			"code":    1,
			"message": "agent:" + e.Error(),
		})
	}
	AllChatDatas[newChat.ChatID] = &newChat
	return ctx.JSON(fiber.Map{
		"code":    0,
		"message": "",
		"chatID":  newChat.ChatID,
	})
}
