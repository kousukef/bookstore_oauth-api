package db

import (
	"github.com/kousukef/bookstore_oauth-api/src/domain/access_token"
	"github.com/kousukef/bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database is not connected yet")
}
