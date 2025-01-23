package models

import "time"

type AuthLog struct {
	ID        string    `json:"id,omitempty"`
	UserID    string    `json:"user_id,omitempty"`
	IPAddress string    `json:"ip_address,omitempty"`
	Action    string    `json:"action,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"` // Server timestamp field
	UpdatedAt time.Time `json:"created_at,omitempty"` // Server timestamp field
}
