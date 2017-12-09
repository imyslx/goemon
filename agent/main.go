package main

import (
	"github.com/imyslx/goemon/agent/item"
	zlog "github.com/rs/zerolog/log"
)

func main() {

	zlog.Info().Msg("Starting Agent.")

	// CreateConfig and get channel for renew.
	ms, _ := item.GetConfig()

	ms.RunAllWokers()

	for {

	}
}
