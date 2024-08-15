package handlers

import (
	"access-point/db"
	"encoding/json"
	"log/slog"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetUserRepo().GetAllUsers()
	if err != nil {
		slog.Error("Error fetching users", "err", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		slog.Error("Error encoding response", "err", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}