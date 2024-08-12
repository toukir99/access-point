package handlers

import (
	"access-point/db"
	"access-point/web/model"
	"access-point/web/utils"
	"encoding/json"
	"log"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	if err := utils.Validate.Struct(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	user.HashPassword()

	query := db.GetQueryBuilder().
			Insert("users").
			Columns("username","email","password").
			Values(user.Username, user.Email, user.Password)
	
	// Convert to SQL and arguments
	sql, args, err := query.ToSql()
	if err != nil {
		log.Fatalf("Error building SQL: %v", err)
	}

	// Execute the query with the write database connection
	_, err = db.WriteDb.Exec(sql, args...)

	//_, err := db.GetWriteDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string {"message": "User created successfully!"})
}