package access_token

import (
	"strings"
	"time"

	"github.com/kousukef/bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string 	`json:"access_token" binding:"required"`
	UserId 			int64 	`json:"user_id" binding:"required"`
	ClientId 		int64 	`json:"client_id" binding:"required"`
	Expires 		int64 	`json:"expires" binding:"required"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
