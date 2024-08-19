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

	userId, err := db.GetUserRepo().CreateUser(&user)
	if err != nil {
		slog.Error("Error creating user", "err", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := SendEmail(user.Email); err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	// response := map[string]interface{}{
	// 	"message": "Email has been sent. Please check your mail and verify the OTP!",
	// 	"user_id": userId,
	// }
	utils.SendData(w, userId)
}
