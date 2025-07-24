package chat

import (
	"mystanford/config"
	"mystanford/person"

	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
)

type Chat struct {
	ChatID          string
	Model           config.ModelConfig
	Persons         []*person.Person
	HistoryMessages []*schema.Message
	Agent           *react.Agent
}
