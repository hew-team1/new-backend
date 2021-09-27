package repository

import (
	"github.com/jmoiron/sqlx"

	"guild-hack-api/app/domain/model"
)

func InsertUser(db *sqlx.Tx, user *model.User) (int64, error) {
	stmt, err := db.Preparex("INSERT INTO users (uid, name, email) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	var id int64
	result, err := stmt.Exec(user.Uid, user.Name, user.Email)
	// この辺りでDBのエラーを
	if err != nil {
		return 0, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func FindUserById(db *sqlx.DB, id int64) (*model.User, error) {
	var user model.User
	err := db.Get(&user, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func AllUser(db *sqlx.DB) ([]model.User, error) {
	var users []model.User
	err := db.Select(&users, "SELECT id, uid, name, email FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func FindUserByUsername(db *sqlx.DB, username string) (*model.User, error) {
	var user model.User
	err := db.Get(&user, "SELECT * FROM users WHERE name = ?", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
