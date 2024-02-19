package storage

import (
	"database/sql"
	"fmt"
	"github.com/zelinskayas/GoBasic/goodRestApiWithDBAuth/internal/app/models"
	"log"
)

// instance of user repository (model inteface)
type UserRepository struct {
	storage *Storage
}

var (
	tableUser string = "dbo.users"
)

// create user in db
func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("USE MyLocalDB INSERT INTO %s (login, password) VALUES(@login, @password) SELECT SCOPE_IDENTITY()", tableUser)
	if err := ur.storage.db.QueryRow(query, sql.Named("login", u.Login), sql.Named("password", u.Password)).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

// find user by login in db
func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	users, err := ur.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var userFinded *models.User
	for _, u := range users {
		if u.Login == login {
			userFinded = u
			founded = true
			break
		}
	}
	return userFinded, founded, nil
}

// select all users in db
func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("USE MyLocalDB SELECT id, login, password FROM %s", tableUser)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//подготовим куда будем читать
	users := make([]*models.User, 0)
	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.ID, &u.Login, &u.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return users, nil
}
