package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"access-point/db"
	"log/slog"
)

type DeleteUserResponse struct {
	Message string `json:"message"`
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
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
	if err := db.GetUserRepo().DeleteUser(id); err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	response := DeleteUserResponse{
		Message: "User deleted successfully",
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("Error encoding response", "err", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
