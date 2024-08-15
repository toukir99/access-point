package handlers

import (
	"access-point/db"
	"access-point/web/model"
	"access-point/web/utils"
	"encoding/json"
	"log/slog"
	"net/http"
)

func SignUpUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := user.HashPassword(); err != nil {
		slog.Error("Error hashing password", "err", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	newID, err := db.GetUserRepo().CreateUser(&user)
	if err != nil {
		slog.Error("Error creating user", "err", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "User created successfully!",
		"user_id": newID,
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("Error encoding response", "err", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
