package models

import (
	"time"
)

type Client struct {
	ID        uint64    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	IsActive  bool      `db:"is_active" json:"is_active"`
	Address   string    `db:"address" json:"address"`
	Currency  string    `db:"currency" json:"currency"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type HarvestClientWrapper struct {
	Clients      []Client `json:"clients"`
	PerPage      int      `json:"per_page"`
	TotalPages   int      `json:"total_pages"`
	TotalEntries int      `json:"total_entries"`
	NextPage     int      `json:"next_page"`
	PreviousPage int      `json:"previous_page"`
	Page         int      `json:"page"`
}
