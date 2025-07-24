package chat

import (
	"context"
	"encoding/json"
	"mystanford/logger"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type EmotionTool struct {
	ChatID string
}

type EmotionModel struct {
	Name string `json:"name"`
	Num  int    `json:"num"`
}

func (t *EmotionTool) Info(_ context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "emotion_change",
		Desc: "更改人物的忍耐值，忍耐值超过一百将会令该人物立刻反驳上一个人的观点",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"name": {
				Type:     schema.String,
				Desc:     "人物的名字",
				Required: true,
			},
			"num": {
				Type:     schema.Integer,
				Desc:     "更改的数值，可以是负数，单次绝对值不超过100",
				Required: true,
			},
		}),
	}, nil
}

func (t *EmotionTool) InvokableRun(_ context.Context, argumentsInJSON string, _ ...tool.Option) (string, error) {
	var data EmotionModel
	e := json.Unmarshal([]byte(argumentsInJSON), &data)
	if e != nil {
		logger.Logger.Warn(t.ChatID, "修改情绪数据:", e.Error())
		return `{"msg": "非法json数据"}`, nil
	}
	nowChat := AllChatDatas[t.ChatID]
	for _, item := range nowChat.Persons {
		if item.Name == data.Name {
			logger.Logger.Debug("修改情绪数据", data.Name, data.Num)
			item.Emotion += data.Num
			break
		}
	}
	return `{"msg": "ok"}`, nil
}
