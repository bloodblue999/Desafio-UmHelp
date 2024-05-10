package main

import (
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/logger"
	"github.com/bloodblue999/umhelp/repo"
	"github.com/bloodblue999/umhelp/server"
	"github.com/bloodblue999/umhelp/server/controller"
	"github.com/bloodblue999/umhelp/service"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	cfg := config.Get()
	logger := logger.New(cfg)

	repo, err := repo.New(cfg)
	if err != nil {
		end(logger, err, "failed to initialize repo manager")
	}

	svc := service.New(cfg, logger, repo)
	ctrl := controller.New(svc, logger)

	svr := server.New(cfg, logger, ctrl)

	if err := svr.Start(); err != nil {
		end(logger, err, "failed to start server")
	}
}

func end(logger *zerolog.Logger, err error, message string) {
	logger.Fatal().Err(err).Msgf("%+v: %+v", message, err)
	time.Sleep(time.Millisecond * 50)

	os.Exit(1)
}
