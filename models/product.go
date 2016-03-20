package models

import (
	"time"
)

type Product struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Provider  string    `json:"provider"`
	Rating    float32   `json:"rating"`
	Status    string    `json:"status"`
	Image     string    `json:"image"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
