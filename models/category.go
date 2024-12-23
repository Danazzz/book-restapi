package models

type Category struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CreatedAt  string `json:"created_at"`
	CreatedBy  string `json:"created_by"`
	ModifiedAt string `json:"modified_at"`
	ModifiedBy string `json:"modified_by"`
}