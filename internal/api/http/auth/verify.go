package auth

import (
	"context"
	resp "my-rest-api/internal/api/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		auth := c.GetHeader("Authentication")

		id, err := h.Service.VerifyToken(ctx, auth)
		if err != nil {
			c.JSON(http.StatusUnauthorized, resp.Response{
				Message: "Unauthenticated",
			})
		}

		ctx = context.WithValue(ctx, "id", id)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
