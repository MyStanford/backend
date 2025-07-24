package server

import (
	"mystanford/server/chat"
	"mystanford/server/model"
	"mystanford/server/person"
)

func AddRoutes() {
	apiGroup := Server.Group("/api")

	personGroup := apiGroup.Group("/person")
	personGroup.Post("/add", person.PersonAddRoute)
	personGroup.Get("/list", person.PersonListRoute)

	chatGroup := apiGroup.Group("/chat")
	chatGroup.Post("/new", chat.NewChatRoute)
	chatGroup.Get("/emotion", chat.EmotionGetRoute)
	chatGroup.Get("/chat", chat.ChatRoute)

	modelGroup := apiGroup.Group("/model")
	modelGroup.Get("/list", model.ModelListRoute)
}
