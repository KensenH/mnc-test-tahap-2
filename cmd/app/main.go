package main

import (
	"context"
	"my-rest-api/internal/boot"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	err := boot.Start(ctx)
	if err != nil {
		logrus.Fatalln(err)
	}
}
