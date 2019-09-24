package manager

import (
	"database/sql"
	"github.com/japhethca/postit-api/db/models"
)

type UserManager struct {
	DB *sql.DB
}

// GetUserByEmail returns a user given an email address
func (um *UserManager) GetUserByEmail(email string) (models.User, error) {
	return um.getUserBy("email", email)
}

// GetUserByID returns a user given a user ID
func (um *UserManager) GetUserByID(userID int) (models.User, error) {
	return um.getUserBy("user_id", userID)
}

func (um *UserManager) getUserBy(field string, value interface{}) (models.User, error) {
	queryStr := "SELECT * FROM ps_user WHERE $1 = $2"
	row := um.DB.QueryRow(queryStr, field, value)
	var user models.User
	err := row.Scan(&user.UserID, &user.Firstname, &user.Lastname, &user.Email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// CreateUser creates a new user in the database
func (um *UserManager) CreateUser(user models.User) (models.User, error) {
	queryStr := `
		Insert Into ps_user (firstname, lastname, email) 
		Values ($1, $2, $3) RETURNING user_id, firstname, lastname, email 
	`
	row := um.DB.QueryRow(queryStr, user.Firstname, user.Lastname, user.Email)
	var newUser models.User
	err := row.Scan(&newUser.UserID, &newUser.Firstname, &newUser.Lastname, &newUser.Email)
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}
