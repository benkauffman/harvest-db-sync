package models

type PartialClient struct {
	ID       uint64 `db:"client_id" json:"id"`
	Name     string `json:"name"`
	Currency string `json:"currency"`
}
