package models

import "time"

type Invoice struct {
	ID             uint64        `db:"id" json:"id"`
	Client         PartialClient `json:"client"`
	Number         string        `db:"number" json:"number"`
	PurchaseOrder  string        `db:"purchase_order" json:"purchase_order"`
	Amount         float64       `db:"amount" json:"amount"`
	DueAmount      float64       `db:"due_amount" json:"due_amount"`
	Tax            float64       `db:"tax" json:"tax"`
	TaxAmount      float64       `db:"tax_amount" json:"tax_amount"`
	Tax2           float64       `db:"tax2" json:"tax2"`
	Tax2Amount     float64       `db:"tax2_amount" json:"tax2_amount"`
	Discount       float64       `db:"discount" json:"discount"`
	DiscountAmount float64       `db:"discount_amount" json:"discount_amount"`
	Subject        string        `db:"subject" json:"subject"`
	Notes          string        `db:"notes" json:"notes"`
	Currency       string        `db:"currency" json:"currency"`
	PeriodStart    string        `db:"period_start,omitempty" json:"period_start"`
	PeriodEnd      string        `db:"period_end,omitempty" json:"period_end"`
	IssueDate      string        `db:"issue_date,omitempty" json:"issue_date"`
	DueDate        string        `db:"due_date,omitempty" json:"due_date"`
	SentAt         time.Time     `db:"sent_at,omitempty" json:"sent_at"`
	PaidAt         time.Time     `db:"paid_at,omitempty" json:"paid_at"`
	ClosedAt       time.Time     `db:"closed_at,omitempty" json:"closed_at"`
	CreatedAt      time.Time     `db:"created_at" json:"created_at"`
	UpdateAt       time.Time     `db:"updated_at" json:"updated_at"`
}

type HarvestInvoiceWrapper struct {
	Invoices     []Invoice `json:"invoices"`
	PerPage      int       `json:"per_page"`
	TotalPages   int       `json:"total_pages"`
	TotalEntries int       `json:"total_entries"`
	NextPage     int       `json:"next_page"`
	PreviousPage int       `json:"previous_page"`
	Page         int       `json:"page"`
}
