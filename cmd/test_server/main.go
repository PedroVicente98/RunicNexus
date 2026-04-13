// Package main is the entry point for the RunicNexus test MMO server.
// It demonstrates how to bootstrap the framework engine and attach modules.
package main

import (
	"log"

	"github.com/PedroVicente98/RunicNexus/pkg/runic"
)

func main() {
	engine := runic.NewEngine(runic.Config{
		Addr: ":7000",
	})

	if err := engine.Start(); err != nil {
		log.Fatalf("engine failed to start: %v", err)
	}
}
