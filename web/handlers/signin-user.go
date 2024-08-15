package handlers

import (
	"access-point/config"
	"access-point/db"
	"access-point/web/model"
	"access-point/web/utils"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt"
)

func SignInUser(w http.ResponseWriter, r *http.Request) {
	var creds model.User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	user, err := db.GetUserRepo().GetUserByEmail(creds.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if err := user.CheckPassword(creds.Password); err != nil {
		slog.Error("Error Invalid password", "err", err)
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &model.Claims { 
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetConfig().JwtSecret))
	if err != nil {
		slog.Error("Error generating token", "err", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "You are logged in",
		"user_id": user.ID,
		"token":   tokenString,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("Error encoding response", "err", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}