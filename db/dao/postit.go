package dao

import (
	"database/sql"
	"github.com/japhethca/postit-api/models"
)

// NewPostit creates a new postit
func NewPostit(db *sql.DB, pt models.Postit) (models.Postit, error) {
	queryStr := `INSERT INTO postit (content, color, user_id, updated_at, created_at) VALUES 
		($1, $2, $3, $4, $5) RETURNING id, content, color, user_id, category_id, updated_at, created_at`

	row := db.QueryRow(queryStr, pt.Content, pt.Color, pt.UserID, pt.UpdatedAt, pt.CreatedAt)
	var newPt models.Postit
	if err := row.Scan(&newPt.ID, &newPt.Content, &newPt.Color, &newPt.UserID, &newPt.CategoryID, &newPt.CreatedAt, &newPt.UpdatedAt); err != nil {
		return models.Postit{}, err
	}
	return newPt, nil
}
