package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/leanovate/microtools/logging"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/untoldwind/amintk/gtk"
	"github.com/untoldwind/trustless-gtk3/state"
	"github.com/untoldwind/trustless-gtk3/ui"
	"github.com/untoldwind/trustless/config"
	"github.com/untoldwind/trustless/daemon"
	"github.com/untoldwind/trustless/secrets"
	"github.com/untoldwind/trustless/secrets/pgp"
	"github.com/untoldwind/trustless/secrets/remote"
)

var rootCommand = &cobra.Command{
	Use:   "Trustless GTK3",
	Short: "GTK3 frontend of trustless",
	Run:   run,
}

var cmdSettings config.Settings
var cfgFile string
var startDaemon bool

func showError(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())

	if cmdSettings.Debug {
		fmt.Fprintln(os.Stderr)
		spew.Fdump(os.Stderr, err)
	}
	os.Exit(1)
}

func init() {
	cobra.MousetrapHelpText = ""
	cobra.OnInitialize(initConfig)

	rootCommand.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default ./trustless.toml, $HOME/.trustless/trustless.toml")
	rootCommand.PersistentFlags().String("node-id", "", "ID of this node")
	rootCommand.PersistentFlags().String("store-url", "", "URL of the store to open (only file:// URLs are supported)")
	rootCommand.PersistentFlags().String("store-scheme", "openpgp", "Storage scheme (openpgp, openpgp+scrypt)")
	rootCommand.PersistentFlags().String("log-file", "", "Log to file instead stdout")
	rootCommand.PersistentFlags().String("log-format", "text", "Log format to use (test, json, logstash)")
	rootCommand.PersistentFlags().Bool("debug", false, "Enable debug information")
	rootCommand.PersistentFlags().Duration("unlock-timeout", 5*time.Minute, "AUtomatic lock timeout")
	rootCommand.PersistentFlags().Bool("unlock-timeout-hard", false, "Enable hard timeout")
	rootCommand.PersistentFlags().BoolVar(&startDaemon, "daemon", false, "Start daemon")

	viper.BindPFlags(rootCommand.PersistentFlags())
}

func initConfig() {
	viper.SetEnvPrefix("TRUSTLESS")
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&cmdSettings); err != nil {

		showError(err)
	}

	if cmdSettings.Debug {
		fmt.Fprintln(os.Stderr, "---- Settings [snip]----")
		spew.Fdump(os.Stderr, cmdSettings)
		fmt.Fprintln(os.Stderr, "---- Settings [snap] ----")
	}

}

func createLogger() logging.Logger {
	loggingOptions := logging.Options{
		Backend:   "simple",
		LogFile:   cmdSettings.LogFile,
		LogFormat: cmdSettings.LogFormat,
		Level:     logging.Info,
	}
	if cmdSettings.Debug {
		loggingOptions.Level = logging.Debug
	}
	return logging.NewLogger(loggingOptions).
		WithContext(map[string]interface{}{"process": "trustless-gtk3"})
}

func readConfig() (*config.Settings, error) {
	var settings config.Settings

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			showError(err)
		}
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.AddConfigPath(filepath.Join(home, ".config"))
		viper.SetConfigName("trustless")
	}

	if err := viper.ReadInConfig(); err != nil {
		if err := writeDefaultConfig(); err != nil {
			return nil, err
		}
	}

	if err := viper.Unmarshal(&settings); err != nil {
		return nil, err
	}

	if settings.Debug {
		fmt.Fprintln(os.Stderr, "---- Settings [snip]----")
		spew.Fdump(os.Stderr, settings)
		fmt.Fprintln(os.Stderr, "---- Settings [snap] ----")
	}

	return &settings, nil
}

func writeDefaultConfig() error {
	if viper.GetString("store-url") == "" {
		viper.Set("store-url", config.DefaultStoreURL())
	}
	if viper.GetString("node-id") == "" {
		nodeId, err := config.GenerateNodeID()
		if err != nil {
			return err
		}
		viper.Set("node-id", nodeId)
	}
	home, err := homedir.Dir()
	if err != nil {
		showError(err)
	}
	os.MkdirAll(filepath.Join(home, ".config"), 0700)

	return viper.WriteConfigAs(filepath.Join(home, ".config", "trustless.toml"))
}

func run(cmd *cobra.Command, args []string) {
	gtk.Init(nil)

	logger := createLogger()
	var secrets secrets.Secrets

	if remote.RemoteAvailable(logger) {
		logger.Info("Use remote store")
		secrets = remote.NewRemoteSecrets(logger)
	} else {
		logger.Info("Starting own daemon")
		settings, err := readConfig()
		if err != nil {
			showError(err)
		}
		scrypted := false
		switch settings.StoreScheme {
		case "openpgp+scrypt":
			scrypted = true
		}

		secrets, err = pgp.NewPGPSecrets(settings.StoreURL, scrypted, settings.NodeID, 4096, settings.UnlockTimeout, settings.UnlockTimeoutHard, logger)
		if err != nil {
			showError(err)
		}
		defer secrets.Lock(context.Background())

		if startDaemon {
			daemon := daemon.NewDaemon(secrets, logger)

			if err := daemon.Start(); err != nil {
				showError(err)
			}
			defer daemon.Stop()
		}
	}

	store, err := state.NewStore(secrets, logger)
	if err != nil {
		logger.ErrorErr(err)
		showError(err)
	}

	window, err := ui.NewMainWindow(store, logger)
	if err != nil {
		logger.ErrorErr(err)
		showError(err)
	}
	window.ShowAll()

	gtk.Main()
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		showError(err)
	}
}
