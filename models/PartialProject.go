package models

type PartialProject struct {
	ID   uint64 `db:"project_id" json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
