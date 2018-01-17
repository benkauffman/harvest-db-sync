package models

import "time"

type User struct {
	ID                           uint64    `db:"id" json:"id"`
	FirstName                    string    `db:"first_name" json:"first_name"`
	LastName                     string    `db:"last_name" json:"last_name"`
	Email                        string    `db:"email" json:"email"`
	Telephone                    string    `db:"telephone" json:"telephone"`
	Timezone                     string    `db:"timezone" json:"timezone"`
	HasAccessToAllFutureProjects bool      `db:"has_access_to_all_future_projects" json:"has_access_to_all_future_projects"`
	IsContractor                 bool      `db:"is_contractor" json:"is_contractor"`
	IsAdmin                      bool      `db:"is_admin" json:"is_admin"`
	IsProjectManager             bool      `db:"is_project_manager" json:"is_project_manager"`
	CanSeeRates                  bool      `db:"can_see_rates" json:"can_see_rates"`
	CanCreateProjects            bool      `db:"can_create_projects" json:"can_create_projects"`
	CanCreateInvoices            bool      `db:"can_create_invoices" json:"can_create_invoices"`
	IsActive                     bool      `db:"is_active" json:"is_active"`
	WeeklyCapacity               uint64    `db:"weekly_capacity" json:"weekly_capacity"`
	DefaultHourlyRate            *float64  `db:"default_hourly_rate" json:"default_hourly_rate"`
	CostRate                     float64   `db:"cost_rate" json:"cost_rate"`
	AvatarUrl                    string    `db:"avatar_URL" json:"avatar_URL"`
	CreatedAt                    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt                    time.Time `db:"updated_at" json:"updated_at"`
}

type HarvestUserWrapper struct {
	Users        []User `json:"users"`
	PerPage      int    `json:"per_page"`
	TotalPages   int    `json:"total_pages"`
	TotalEntries int    `json:"total_entries"`
	NextPage     int    `json:"next_page"`
	PreviousPage int    `json:"previous_page"`
	Page         int    `json:"page"`
}
