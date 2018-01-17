package models

type PartialInvoice struct {
	ID     uint64 `db:"invoice_id,omitempty" json:"id"`
	Number string `json:"number"`
}
