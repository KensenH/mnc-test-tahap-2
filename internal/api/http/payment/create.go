package payment

import (
	resp "my-rest-api/internal/api/http"
	"my-rest-api/internal/domain/payment"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (h *Handler) Create(c *gin.Context) {
	var (
		req payment.CreatePaymentReq
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

	id := ctx.Value("id").(string)

	if id == "" {
		c.JSON(http.StatusBadRequest, resp.Response{
			Message: "Unauthenticated",
		})
		return
	}

	jwt, err := h.Service.Create(ctx, id, req.Amount, req.Remarks)
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
