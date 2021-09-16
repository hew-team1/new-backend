package model

import (
	"time"

	"github.com/go-openapi/strfmt"

	"guild-hack-api/swagger/generated_swagger"
)

type User struct {
	Id        int64     `db:"id"`
	Uid       string    `db:"uid"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (u *User) CreateResponseSwaggerModel() *generated_swagger.CreateUserResponse {
	return &generated_swagger.CreateUserResponse{
		ID:        u.Id,
		UID:       u.Uid,
		Name:      u.Name,
		Email:     strfmt.Email(u.Email),
		CreatedAt: u.CreatedAt.Format("2006-01-02T15:04"),
		UpdatedAt: u.UpdatedAt.Format("2006-01-02T15:04"),
	}
}
