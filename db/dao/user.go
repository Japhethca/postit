package dao

import (
	"database/sql"
	"github.com/japhethca/postit-api/models"
)

type UserDAO struct {
	DB *sql.DB
}

// GetUserByEmail returns a user given an email address
func (ud *UserDAO) GetUserByEmail(email string) (models.User, error) {
	queryStr := "SELECT * FROM ps_user WHERE email = $1"
	return ud.getUser(queryStr, email)
}

// GetUserByID returns a user given a user ID
func (ud *UserDAO) GetUserByID(userID int) (models.User, error) {
	queryStr := "SELECT * FROM ps_user WHERE user_id = $1"
	return ud.getUser(queryStr, userID)
}

func (ud *UserDAO) getUser(qs string, values ...interface{}) (models.User, error) {
	row := ud.DB.QueryRow(qs, values...)
	var user models.User
	err := row.Scan(&user.UserID, &user.Firstname, &user.Lastname, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// CreateUser creates a new user in the database
func (ud *UserDAO) CreateUser(user models.User) (models.User, error) {
	queryStr := `
		Insert Into ps_user (firstname, lastname, email, password) 
		Values ($1, $2, $3, $4) RETURNING user_id, firstname, lastname, email 
	`
	row := ud.DB.QueryRow(queryStr, user.Firstname, user.Lastname, user.Email, user.Password)
	var newUser models.User
	err := row.Scan(&newUser.UserID, &newUser.Firstname, &newUser.Lastname, &newUser.Email, &newUser.Password)
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}
