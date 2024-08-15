package handlers

import (
	"access-point/db"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	user, err := db.GetUserRepo().GetUserByID(id)
	if err != nil {
		http.Error(w, "Error retrieving user info", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		slog.Error("Error encoding response", "err", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}