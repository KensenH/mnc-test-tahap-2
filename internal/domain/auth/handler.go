package auth

import "github.com/gin-gonic/gin"

type Handler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Verify() gin.HandlerFunc
}
