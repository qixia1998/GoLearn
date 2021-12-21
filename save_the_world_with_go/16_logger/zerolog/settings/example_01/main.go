package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
)

func main() {
	tempFile, err := ioutil.TempFile(os.TempDir(), "deleteme")
	if err != nil {
		log.Error().Err(err).Msg("there was an error creating a temporary file four our log")
	}
	defer tempFile.Close()
	fileLogger := zerolog.New(tempFile).With().Logger()
	fileLogger.Info().Msg("This is an entry from my log")
	fmt.Printf("The log file is allocated at %s\n", tempFile.Name())
}
