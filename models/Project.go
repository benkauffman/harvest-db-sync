package models

import "time"

type Project struct {
	ID                               uint64        `db:"id" json:"id"`
	ProjectClient                    PartialClient `json:"client"`
	Name                             string        `db:"name" json:"name"`
	Code                             string        `db:"code" json:"code"`
	IsActive                         bool          `db:"is_active" json:"is_active"`
	IsBillable                       bool          `db:"is_billable" json:"is_billable"`
	IsFixedFee                       bool          `db:"is_fixed_fee" json:"is_fixed_fee"`
	BillBy                           string        `db:"bill_by" json:"bill_by"`
	HourlyRate                       float64       `db:"hourly_rate" json:"hourly_rate"`
	Budget                           *float64      `db:"budget" json:"budget"`
	BudgetBy                         string        `db:"budget_by" json:"budget_by"`
	NotifyWhenOverBudget             bool          `db:"notify_when_over_budget" json:"notify_when_over_budget"`
	OverBudgetNotificationPercentage float64       `db:"over_budget_notification_percentage" json:"over_budget_notification_percentage"`
	OverBudgetNotificationDate       string        `db:"over_budget_notification_date,omitempty" json:"over_budget_notification_date"`
	ShowBudgetToAll                  bool          `db:"show_budget_to_all" json:"show_budget_to_all"`
	CostBudget                       *float64      `db:"cost_budget" json:"cost_budget"`
	CostBudgetIncludeExpense         bool          `db:"cost_budget_include_expenses" json:"cost_budget_include_expenses"`
	Fee                              *float64      `db:"fee" json:"fee"`
	Notes                            string        `db:"notes" json:"notes"`
	StartsOn                         string        `db:"starts_on,omitempty" json:"starts_on"`
	EndsOn                           string        `db:"ends_on,omitempty" json:"ends_on"`
	CreatedAt                        time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt                        time.Time     `db:"updated_at" json:"updated_at"`
}

type HarvestProjectWrapper struct {
	Projects     []Project `json:"projects"`
	PerPage      int       `json:"per_page"`
	TotalPages   int       `json:"total_pages"`
	TotalEntries int       `json:"total_entries"`
	NextPage     int       `json:"next_page"`
	PreviousPage int       `json:"previous_page"`
	Page         int       `json:"page"`
}
