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
	if err != nil {
		return 0, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func FindUser(db *sqlx.DB, id int64) (*model.User, error) {
	var user model.User
	err := db.Get(&user, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
