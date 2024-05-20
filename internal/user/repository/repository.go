package repository

import (
	"eulabs/products/internal/user"
	"eulabs/products/internal/user/dtos"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) user.Repository {
	return &repository{db: db}
}

func (r *repository) GetUserByLogin(login string) (result *dtos.UserLoginRequest, err error) {

	row := r.db.QueryRowx(`SELECT login, password FROM user WHERE login=?`, login)
	err = row.Err()
	if err != nil {
		return
	}

	result = &dtos.UserLoginRequest{}
	row.StructScan(result)

	return
}
