package chat

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudwego/eino/schema"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type ChatMessageModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Role    string `json:"role"`
	Content string `json:"content"`
}

func SendChatSSEMessage(w *bufio.Writer, message ChatMessageModel) {
	m, _ := json.Marshal(message)
	fmt.Fprintf(w, "data: %s\n\n", m)
	w.Flush()
}

func ChatRoute(ctx *fiber.Ctx) error {
	chatID := ctx.Query("chatID", "")
	message := ctx.Query("message", "")

	ctx.Set("Content-Type", "text/event-stream")
	ctx.Set("Cache-Control", "no-cache")
	ctx.Set("Connection", "keep-alive")
	ctx.Set("Transfer-Encoding", "chunked")

	ctx.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		if chatID == "" {
			SendChatSSEMessage(w, ChatMessageModel{
				Code:    1,
				Message: "chatID不能为空",
			})
			return
		}
		if message == "" {
			SendChatSSEMessage(w, ChatMessageModel{
				Code:    1,
				Message: "message不能为空",
			})
			return
		}
		nowChat, ok := AllChatDatas[chatID]
		if !ok {
			SendChatSSEMessage(w, ChatMessageModel{
				Code:    1,
				Message: "chatID不存在",
			})
			return
		}
		for _, item := range nowChat.Persons {
			nowChat.HistoryMessages = append(nowChat.HistoryMessages, &schema.Message{
				Role:    "user",
				Content: message + ";" + item.Name,
			})
			res, e := nowChat.Agent.Generate(context.Background(), nowChat.HistoryMessages)
			if e != nil {
				SendChatSSEMessage(w, ChatMessageModel{
					Code:    2,
					Message: e.Error(),
				})
				continue
			}
			nowChat.HistoryMessages = append(nowChat.HistoryMessages, &schema.Message{
				Role:    schema.Assistant,
				Content: res.Content,
			})
			SendChatSSEMessage(w, ChatMessageModel{
				Role:    item.Name,
				Content: res.Content,
			})
		}
	}))

	return nil
}
