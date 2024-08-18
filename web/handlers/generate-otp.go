package handlers

import (
	"access-point/db"
	"context"
	"fmt"
	"net/smtp"
	"time"

	"github.com/google/uuid"
)

func SendEmail(userEmail string) error{

	otp := uuid.NewString()[0:6]

	const (
		otpPrefix       = "otp:"        
		otpExpiration   = 5 * time.Minute 
	)
	
	otpKey := otpPrefix + userEmail

	err := db.GetRedisClient().Set(context.Background(), otpKey, otp, otpExpiration).Err()
	if err != nil {
		return fmt.Errorf("failed to store OTP: %v", err)
	}

	from := "mailto:mdfazlerabbispondontechnonext@gmail.com"
	password := "anyu uzkc hpaw hzcl"

	to := []string{userEmail}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Subject: Your OTP Code\n\n" +
		"Your OTP code is: " + otp)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return fmt.Errorf("error sending email")
	}

	return nil
}
