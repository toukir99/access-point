package logger

import "log/slog"


var Error = slog.Error

type LogKey string

const ExtraKey  LogKey = "extra"

func Extra(value any) slog.Attr {
	return slog.String(string(ExtraKey), ConvertToJson(value))
}