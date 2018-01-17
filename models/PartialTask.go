package models

type PartialTask struct {
	ID   uint64 `db:"task_id" json:"id"`
	Name string `json:"name"`
}
