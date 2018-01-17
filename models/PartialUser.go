package models

type PartialUser struct {
	ID   uint64 `db:"user_id" json:"id"`
	Name string `json:"name"`
}
