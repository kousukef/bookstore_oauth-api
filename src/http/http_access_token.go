package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kousukef/bookstore_oauth-api/src/domain/access_token"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := h.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}
