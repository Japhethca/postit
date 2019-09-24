package db

import "database/sql"

const (
	// DatabaseUniqueViolation is the code for postgres database unique violation error
	DatabaseUniqueViolation     = "23505"
	DatabaseForeignKeyViolation = "23503"
	DatabaseNotNullViolation    = "23502"
)

// Setter provides an interface of method for setting databases to objects
type Setter interface {
	SetDB(*sql.DB)
}
