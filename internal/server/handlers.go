package server

func (s Server) registerHandlers() {
	s.Router.POST("/register", s.AuthHandler.Register)
	s.Router.POST("/login", s.AuthHandler.Login)

	s.Router.Use(s.AuthHandler.Verify()).POST("/topup", s.TopupHandler.Create)
	s.Router.Use(s.AuthHandler.Verify()).POST("/payment", s.TopupHandler.Create)

	//TO-DO: transfer
}
