package topup

import "github.com/gin-gonic/gin"

type Handler interface {
	Create(c *gin.Context)
}
