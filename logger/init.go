package logger

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
)

func InitLogger() {
	Logger = slog.New()
	consoleHandler := handler.NewConsoleHandler(slog.AllLevels)
	consoleHandler.Formatter().(*slog.TextFormatter).SetTemplate(LoggerTemplate)
	Logger.AddHandler(consoleHandler)
}
