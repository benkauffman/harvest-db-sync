package models

import "time"

type TimeEntry struct {
	ID             uint64         `db:"id" json:"id"`
	User           PartialUser    `json:"user"`
	Client         PartialClient  `json:"client"`
	Project        PartialProject `json:"project"`
	Task           PartialTask    `json:"task"`
	Invoice        PartialInvoice `json:"invoice"`
	SpentDate      string         `db:"spent_date" json:"spent_date"`
	Hours          float64        `db:"hours" json:"hours"`
	Notes          string         `db:"notes" json:"notes"`
	IsLocked       bool           `db:"is_locked" json:"is_locked"`
	LockedReason   string         `db:"locked_reason" json:"locked_reason"`
	IsClosed       bool           `db:"is_closed" json:"is_closed"`
	IsBilled       bool           `db:"is_billed" json:"is_billed"`
	TimerStartedAt string         `db:"timer_started_at,omitempty" json:"timer_started_at"`
	StartedTime    string         `db:"started_time,omitempty" json:"started_time"`
	EndedTime      string         `db:"ended_time,omitempty" json:"ended_time"`
	IsRunning      bool           `db:"is_running" json:"is_running"`
	Billable       bool           `db:"billable" json:"billable"`
	Budgeted       bool           `db:"budgeted" json:"budgeted"`
	BillableRate   *float64       `db:"billable_rate" json:"billable_rate"`
	CostRate       *float64       `db:"cost_rate" json:"cost_rate"`
	CreatedAt      time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at" json:"updated_at"`
}

type HarvestTimeEntryWrapper struct {
	TimeEntries  []TimeEntry `json:"time_entries"`
	PerPage      int         `json:"per_page"`
	TotalPages   int         `json:"total_pages"`
	TotalEntries int         `json:"total_entries"`
	NextPage     int         `json:"next_page"`
	PreviousPage int         `json:"previous_page"`
	Page         int         `json:"page"`
}
