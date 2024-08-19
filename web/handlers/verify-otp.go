package handlers

import (
	"access-point/db"
	"access-point/web/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

func VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var user model.OTPRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if user.Email == "" || user.OTP == "" {
		http.Error(w, "email and otp are required", http.StatusBadRequest)
		return
	}

	otpKey := "otp:" + user.Email

	storedOTP, err := db.GetRedisClient().Get(context.Background(), otpKey).Result()
	if err != nil {
		if err == redis.Nil {
			http.Error(w, "OTP not found or expired", http.StatusUnauthorized)
			return
		}
		http.Error(w, fmt.Sprintf("error fetching OTP: %v", err), http.StatusInternalServerError)
		return
	}

	if storedOTP != user.OTP {
		http.Error(w, "invalid OTP", http.StatusUnauthorized)
		return
	}

	if err := db.GetUserRepo().ActivateUserByEmail(user.Email); err != nil {
		http.Error(w, fmt.Sprintf("Error updating user status: %v", err), http.StatusInternalServerError)
		return
	}

	_, err = db.GetRedisClient().Del(context.Background(), otpKey).Result()
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting OTP: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OTP is verified successfully and you have registered successfully!"))
}
