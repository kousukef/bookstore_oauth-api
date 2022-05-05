package db

import (
	"github.com/gocql/gocql"
	"github.com/kousukef/bookstore_oauth-api/src/clients/cassandras"
	"github.com/kousukef/bookstore_oauth-api/src/domain/access_token"
	"github.com/kousukef/bookstore_oauth-api/src/utils/errors"
)

var (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES (?,?,?,?);"
	queryUpdateExpires = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationtime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	var result access_token.AccessToken
	if err := cassandras.GetSession().Query(queryGetAccessToken, id).Scan(
			&result.AccessToken,
			&result.UserId,
			&result.ClientId,
			&result.Expires,
		); err != nil {
			if err == gocql.ErrNotFound {
				return nil, errors.NewNotFoundError("no access token is found")
			}
			return nil, errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	if err := cassandras.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpirationtime(at access_token.AccessToken) *errors.RestErr {
	if err := cassandras.GetSession().Query(queryUpdateExpires,
			at.Expires,
			at.AccessToken,
		).Exec(); err != nil {
			return errors.NewBadRequestError(err.Error())
	}

	return nil
}
