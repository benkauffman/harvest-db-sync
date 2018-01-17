package models

import "time"

type Task struct {
	ID                uint64    `db:"id" json:"id"`
	Name              string    `db:"name" json:"name"`
	BillableByDefault bool      `db:"billable_by_default" json:"billable_by_default"`
	DefaultHourlyRate float64   `db:"default_hourly_rate" json:"default_hourly_rate"`
	IsDefault         bool      `db:"is_default" json:"is_default"`
	IsActive          bool      `db:"is_active" json:"is_active"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	UpdateAt          time.Time `db:"updated_at" json:"updated_at"`
}

type HarvestTaskWrapper struct {
	Tasks        []Task `json:"tasks"`
	PerPage      int    `json:"per_page"`
	TotalPages   int    `json:"total_pages"`
	TotalEntries int    `json:"total_entries"`
	NextPage     int    `json:"next_page"`
	PreviousPage int    `json:"previous_page"`
	Page         int    `json:"page"`
}
