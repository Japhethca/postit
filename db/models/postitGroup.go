package models

// PostitGroup model
type PostitGroup struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}
