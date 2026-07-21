package main

import (
	"github.com/frontdex/localsandbox/cmd"
	"github.com/frontdex/localsandbox/internal/logging"
)

func main() {
	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("Application terminated due to unhandled panic")
	})

	cmd.Execute()
}
