package logger

import "github.com/gookit/slog"

var (
	Logger *slog.Logger
)

const (
	LoggerTemplate = "[{{datetime}}] [{{level}}] [{{caller}}] {{message}} {{data}} {{extra}}\n"
)
