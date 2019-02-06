package models

import "time"

type LXD struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	IP        string    `json:"ip"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
