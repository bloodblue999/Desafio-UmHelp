package logger

import (
	"github.com/bloodblue999/umhelp/config"
	"os"

	"github.com/rs/zerolog"
)

func New(cfg *config.Config) *zerolog.Logger {
	logger := zerolog.New(os.Stderr).With().Str("service", cfg.InternalConfig.ServiceName).Timestamp().Logger()

	return &logger
}
