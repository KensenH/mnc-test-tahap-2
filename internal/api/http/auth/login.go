package auth

import (
	resp "my-rest-api/internal/api/http"
	"my-rest-api/internal/domain/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (h *Handler) Login(c *gin.Context) {
	var (
		req auth.LoginReq
		ctx = c.Request.Context()
		v   = validator.New()
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Response{
			Message: err.Error(),
		})
		return
	}

	err = v.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Response{
			Message: err.Error(),
		})
		return
	}

	jwt, err := h.Service.Login(ctx, req)
	if err != nil {
		c.JSON(http.StatusOK, resp.Response{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp.Response{
		Status: "SUCCESS",
		Result: jwt,
	})
}
