package main

import (
	"context"
	"os"

	"github.com/gotk3/gotk3/gtk"
	"github.com/leanovate/microtools/logging"
	"github.com/untoldwind/trustless-gtk3/state"
	"github.com/untoldwind/trustless-gtk3/ui"
	"github.com/untoldwind/trustless/config"
	"github.com/untoldwind/trustless/daemon"
	"github.com/untoldwind/trustless/secrets"
	"github.com/untoldwind/trustless/secrets/pgp"
	"github.com/untoldwind/trustless/secrets/remote"
	cli "gopkg.in/urfave/cli.v2"
)

var GlobalFlags = &struct {
	Debug      bool
	LogFormat  string
	LogFile    string
	ConfigFile string
	NoDaemon   bool
}{}

func createLogger() logging.Logger {
	loggingOptions := logging.Options{
		Backend:   "logrus",
		LogFile:   GlobalFlags.LogFile,
		LogFormat: GlobalFlags.LogFormat,
		Level:     logging.Info,
	}
	if GlobalFlags.Debug {
		loggingOptions.Level = logging.Debug
	}
	return logging.NewLogrusLogger(loggingOptions).
		WithContext(map[string]interface{}{"process": "trustless-gtk3"})
}

func main() {
	app := &cli.App{
		Name:  "Trustless GTK3",
		Usage: "GTK3 frontend of trustless",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "log-file",
				Value:       "",
				Usage:       "Log to file instead stdout",
				EnvVars:     []string{"DEPLOY_CONTROL_LOG_FILE"},
				Destination: &GlobalFlags.LogFile,
			},
			&cli.StringFlag{
				Name:        "log-format",
				Value:       "text",
				Usage:       "Log format to use (test, json, logstash)",
				EnvVars:     []string{"DEPLOY_CONTROL_LOG_FORMAT"},
				Destination: &GlobalFlags.LogFormat,
			},
			&cli.BoolFlag{
				Name:        "debug",
				Usage:       "Enable debug logging",
				EnvVars:     []string{"DEPLOY_CONTROL_DEBUG"},
				Destination: &GlobalFlags.Debug,
			},
			&cli.StringFlag{
				Name:        "config-file",
				Usage:       "Client configuration file",
				Value:       config.DefaultConfigFile(),
				Destination: &GlobalFlags.ConfigFile,
			},
			&cli.BoolFlag{
				Name:        "no-daemon",
				Usage:       "Never start embedded daemon",
				Destination: &GlobalFlags.NoDaemon,
			},
		},
		Action: run,
	}

	app.Run(os.Args)
}

func run(ctx *cli.Context) error {
	gtk.Init(nil)

	logger := createLogger()
	var secrets secrets.Secrets

	if GlobalFlags.NoDaemon || remote.RemoteAvailable(logger) {
		logger.Info("Use remote store")
		secrets = remote.NewRemoteSecrets(logger)
	} else {
		logger.Info("Starting own daemon")
		config, err := config.ReadConfig(GlobalFlags.ConfigFile, logger)
		if err != nil {
			return err
		}

		secrets, err = pgp.NewPGPSecrets(config.StoreURL, config.NodeID, 4096, config.UnlockTimeout, config.UnlockTimeoutHard, logger)
		if err != nil {
			return err
		}
		defer secrets.Lock(context.Background())

		daemon := daemon.NewDaemon(secrets, logger)

		if err := daemon.Start(); err != nil {
			return err
		}
		defer daemon.Stop()
	}

	store, err := state.NewStore(secrets, logger)
	if err != nil {
		logger.ErrorErr(err)
		return err
	}

	window, err := ui.NewMainWindow(store, logger)
	if err != nil {
		logger.ErrorErr(err)
		return err
	}
	window.ShowAll()

	gtk.Main()

	return nil
}
