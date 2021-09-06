package model

import "time"

type EmailInformation struct {
	Id        int       `json:"id,omitempty"`
	Sent      bool      `json:"sent,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email,omitempty"`
	Name      string    `json:"name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
}
