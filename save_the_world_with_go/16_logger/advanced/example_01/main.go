package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"math/rand"
)

type ComponentHock struct {
	component string
}

func (h ComponentHock) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level == zerolog.DebugLevel {
		e.Str("component", h.component)
	}
}

type RandomHook struct {}

func (r RandomHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	e.Int("random", rand.Int())
}

func main() {
	logger := log.Hook(ComponentHock{"mouduleA"})
	logger = logger.Hook(RandomHook{})
	logger.Info().Msg("Info message")
	logger.Debug().Msg("Debug message")
}