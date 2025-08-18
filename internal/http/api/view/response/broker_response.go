package response

import "time"

type Broker struct {
	ID        uint       `json:"id"`
	Nome      string     `json:"nome"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
