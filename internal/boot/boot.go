package boot

import (
	"context"
	"my-rest-api/internal/config"
	"my-rest-api/internal/init/db"
	"my-rest-api/internal/server"

	authH "my-rest-api/internal/api/http/auth"
	userD "my-rest-api/internal/data/users"
	authS "my-rest-api/internal/service/auth"

	topupH "my-rest-api/internal/api/http/topup"
	topupD "my-rest-api/internal/data/topup"
	topupS "my-rest-api/internal/service/topup"

	paymentH "my-rest-api/internal/api/http/payment"
	paymentD "my-rest-api/internal/data/payment"
	paymentS "my-rest-api/internal/service/payment"
)

func Start(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	database, err := db.New(ctx, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbHost, cfg.DbPort)
	if err != nil {
		return err
	}

	defer database.Close(ctx)

	userData := userD.New(database.Connection)
	authService := authS.New(userData, cfg.JwtSecret)
	authHandler := authH.New(authService)

	topupData := topupD.New(database.Connection)
	topupService := topupS.New(topupData)
	topupHandler := topupH.New(topupService)

	paymentData := paymentD.New(database.Connection)
	paymentService := paymentS.New(paymentData)
	paymentHandler := paymentH.New(paymentService)

	s := server.New(
		authHandler,
		topupHandler,
		paymentHandler,
		server.WithHttpPort(cfg.HttpPort),
	)

	return s.Start()
}
