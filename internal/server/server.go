package server

import (
	"my-rest-api/internal/domain/auth"
	"my-rest-api/internal/domain/payment"
	"my-rest-api/internal/domain/topup"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Option func(s *Server)

type Server struct {
	HttpPort       string
	Router         *gin.Engine
	AuthHandler    auth.Handler
	TopupHandler   topup.Handler
	PaymentHandler payment.Handler
}

func New(
	authH auth.Handler,
	topupH topup.Handler,
	paymentH payment.Handler,
	opts ...Option,
) *Server {
	r := gin.Default()
	r.Use(gin.Recovery())

	s := Server{
		HttpPort:       ":8080", // default port
		Router:         r,
		AuthHandler:    authH,
		TopupHandler:   topupH,
		PaymentHandler: paymentH,
	}

	r.GET("/", Default)

	for _, o := range opts {
		o(&s)
	}

	s.registerHandlers()

	return &s
}

func (s *Server) Start() error {
	return s.Router.Run(s.HttpPort)
}

func WithHttpPort(port string) Option {
	return func(s *Server) {
		s.HttpPort = port
	}
}

func Default(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
