package adapter

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Zerolog store main configurations for zerolog
type Zerolog struct {
	logger zerolog.Logger
}

// NewZerolog used to create new instance
func NewZerolog() *Zerolog {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger()
	return &Zerolog{log}
}

// Debug .
func (z *Zerolog) Debug(msg string) {
	z.logger.Debug().Msg(msg)
}

// Debugf .
func (z *Zerolog) Debugf(format string, v ...interface{}) {
	z.logger.Debug().Msgf(format, v...)
}

// Warn .
func (z *Zerolog) Warn(msg string) {
	z.logger.Warn().Msg(msg)
}

// Warnf .
func (z *Zerolog) Warnf(format string, v ...interface{}) {
	z.logger.Warn().Msgf(format, v...)
}

// Info .
func (z *Zerolog) Info(msg string) {
	z.logger.Info().Msg(msg)
}

// Infof .
func (z *Zerolog) Infof(format string, v ...interface{}) {
	z.logger.Info().Msgf(format, v...)
}

// Error .
func (z *Zerolog) Error(msg string) {
	z.logger.Error().Msg(msg)
}

// Errorf .
func (z *Zerolog) Errorf(format string, v ...interface{}) {
	z.logger.Error().Msgf(format, v...)
}

// Fatal .
func (z *Zerolog) Fatal(msg string) {
	z.logger.Fatal().Msg(msg)
}

// Fatalf .
func (z *Zerolog) Fatalf(format string, v ...interface{}) {
	z.logger.Fatal().Msgf(format, v...)
}
