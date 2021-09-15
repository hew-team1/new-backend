package usecase

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"guild-hack-api/app/db/dbutil"
	"guild-hack-api/app/domain/model"
	"guild-hack-api/app/domain/repository"
)

type UserUseCase struct {
	db *sqlx.DB
}

func NewUserUseCase(db *sqlx.DB) *UserUseCase {
	return &UserUseCase{
		db: db,
	}
}
func (u *UserUseCase) Create(uid string, name string, email string) (int64, error) {
	newUser := &model.User{
		Uid:   uid,
		Name:  name,
		Email: email,
	}

	var createId int64
	if err := dbutil.TXHandler(u.db, func(tx *sqlx.Tx) error {
		id, err := repository.InsertUser(tx, newUser)
		if err != nil {
			return err
		}
		createId = id
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return 0, fmt.Errorf("failed user insert transaction: %w", err)
	}
	return createId, nil
}

func (u *UserUseCase) Show(id int64) (*model.User, error) {
	return repository.FindUser(u.db, id)
}
