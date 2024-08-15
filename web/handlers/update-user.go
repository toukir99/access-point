package handlers

import (
	"access-point/db"
	"access-point/web/model"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type UpdateUserResponse struct {
	Message string `json:"message"`
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req model.User
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

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

	// Hash the new password
	//hashedPassword := req.HashPassword()

	err = db.GetUserRepo().UpdateUser(id, &req)
	if err != nil {
		http.Error(w, "Error updating user info", http.StatusInternalServerError)
		return
	}

	response := UpdateUserResponse{
		Message: "User updated successfully",
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("Error encoding response", "err", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
