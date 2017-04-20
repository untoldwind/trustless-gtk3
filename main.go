package main

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/trustless-gtk3/state"
	"github.com/untoldwind/trustless-gtk3/ui"
	"github.com/untoldwind/trustless/secrets/remote"
)

func createLogger() logging.Logger {
	loggingOptions := logging.Options{
		Backend:   "logrus",
		LogFormat: "text",
		Level:     logging.Info,
	}
	return logging.NewLogrusLogger(loggingOptions).
		WithContext(map[string]interface{}{"process": "trustless-q5"})
}

func main() {
	gtk.Init(nil)

	logger := createLogger()
	secrets := remote.NewRemoteSecrets(logger)

	store, err := state.NewStore(secrets, logger)
	if err != nil {
		logger.ErrorErr(err)
		return
	}

	window, err := ui.NewMainWindow(store, logger)
	if err != nil {
		logger.ErrorErr(err)
		return
	}
	window.ShowAll()

	gtk.Main()
}
