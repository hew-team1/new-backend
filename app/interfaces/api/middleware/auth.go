package middleware

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/jmoiron/sqlx"

	"guild-hack-api/app/domain/repository"
	"guild-hack-api/app/interfaces/api/httputil"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{
		db: db,
	}
}

func (auth *Auth) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		a := r.Header.Get("Authorization")
		pair := strings.SplitN(a, " ", 2)
		if len(pair) < 2 {
			httputil.RespondErrorJson(w, http.StatusUnauthorized, errors.New("not enough authentication"))
		}

		authType := pair[0]
		if !strings.EqualFold(authType, "username") {
			httputil.RespondErrorJson(w, http.StatusUnauthorized, errors.New("username not found"))
		}

		username := pair[1]
		if user, err := repository.FindUserByUsername(auth.db, username); err != nil {
			if err == sql.ErrNoRows {
				httputil.RespondErrorJson(w, http.StatusUnauthorized, errors.New("the authenticated user did not exist"))
				return
			}
			
			httputil.RespondErrorJson(w, http.StatusInternalServerError, err)
			return
		} else {
			ctx = httputil.SetUserToContext(r.Context(), user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
