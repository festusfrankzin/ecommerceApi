package models

import (

	"database/sql"

)

type RegisterUserPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type User struct {
	db *sql.DB
}
