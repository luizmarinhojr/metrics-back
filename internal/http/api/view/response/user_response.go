package response

import "time"

type User struct {
	ID        uint       `json:"id"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserName struct {
    Corretor string `json:"corretor"`
}