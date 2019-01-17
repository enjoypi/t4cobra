package cmd

import (
	//_ "github.com/spf13/viper/remote"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile           string
	configRemoteEndpoint string
	configRemotePath     string
	configRemoteType     string
	logLevel             string
	rootViper            = viper.New()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "t4cobra",
	Short: "the template of cobra",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application`,

	PreRunE: preRunE,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(rootViper)
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.t4cobra.yaml)")

	rootCmd.PersistentFlags().StringVar(&configRemoteEndpoint,
		"config.remote.endpoint",
		"127.0.0.1:2379",
		"the endpoint of remote config")
	rootCmd.PersistentFlags().StringVar(&configRemoteType,
		"config.remote.type",
		"etcd",
		"config file (default is $HOME/.t4cobra.yaml)")
	rootCmd.PersistentFlags().StringVar(&configRemotePath,
		"config.remote.path",
		"/t4cobra/config",
		"config file (default is $HOME/.t4cobra.yaml)")

	rootCmd.PersistentFlags().StringVar(&logLevel, "log.level", "info", "level of logrus")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().Bool("version", false, "show version")
}

func preRunE(cmd *cobra.Command, args []string) error {
	// Viper uses the following precedence order. Each item takes precedence over the item below it:
	//
	// explicit call to Set
	// flag
	// env
	// config
	// key/value store
	// default
	//
	// Viper configuration keys are case insensitive.

	v := rootViper

	// remote config
	//if err := v.AddRemoteProvider(configRemoteType, configRemoteEndpoint, configRemotePath); err != nil {
	//	return err
	//}
	//viper.SetConfigType("json")
	//if err := v.ReadRemoteConfig(); err != nil {
	//	return err
	//}

	// local config
	if configFile != "" {
		// Use config file from the flag.
		v.SetConfigFile(configFile)

		// If a config file is found, read it in.
		if err := v.ReadInConfig(); err != nil {
			return err
		}
		logrus.Debug("using config file: ", v.ConfigFileUsed())
		logrus.Debug("settings in config file: ", v.AllSettings())
	}

	v.AutomaticEnv() // read in environment variables that match

	ll := v.GetString("log.level")
	if ll != "" {
		logLevel = ll
	}

	if lvl, err := logrus.ParseLevel(logLevel); err == nil {
		logrus.SetLevel(lvl)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.Warn(err)
	}

	if err := v.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	logrus.Info("current log level: ", logrus.GetLevel())
	logrus.Debug("all settings: ", v.AllSettings())
	return nil
}
