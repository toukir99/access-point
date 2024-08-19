package utils

import (
	"access-point/logger"
	"encoding/json"
	"log/slog"
	"net/http"
)

func SendJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")

	str, err := json.Marshal(data)
	if err != nil {
		slog.Error(err.Error(), logger.Extra(map[string]any{
			"data": data,
		}))

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(`{"message": "Internal Server Error!"}`))
		return 
	}

	w.WriteHeader(status)
	w.Write(str)
}