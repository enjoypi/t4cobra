package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	rootViper = viper.New()
)

func run(v *viper.Viper) error {
	logrus.Info("all settings: ", v.AllSettings())
	return fmt.Errorf("some error")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bootcobra",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		//if err := rootViper.BindPFlags(cmd.Flags()); err != nil {
		//	return err
		//}
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
	logrus.Trace("root.init")
	defer func() {
		logrus.Trace("root.init ok")
	}()
		cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bootcobra.yaml)")
	rootCmd.PersistentFlags().String("log.level", "info", "level of logrus")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().Bool("root.flag", true, "flag of root command")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	logrus.Trace("root.initConfig")
	defer func() {
		logrus.Trace("root.initConfig ok")
	}()

	v := rootViper

	if cfgFile != "" {
		// Use config file from the flag.
		v.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".bootcobra" (without extension).
		v.AddConfigPath(home)
		v.SetConfigName(".bootcobra")
	}

	v.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := v.ReadInConfig(); err == nil {
		logrus.Info("Using config file: ", v.ConfigFileUsed())
		logrus.Info("settings in config file: ", v.AllSettings())
	}

	if err:= v.BindPFlags(rootCmd.Flags()); err!=nil {
		logrus.Fatal(err)
	}

	if lvl, err:=logrus.ParseLevel(v.GetString("log.level")); err == nil {
		logrus.SetLevel(lvl)
	}else {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.Warn(err)
	}

	logrus.Info("current log level: ", logrus.GetLevel())
	logrus.Info("global settings:", v.AllSettings())
}
