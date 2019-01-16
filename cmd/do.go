package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//if err := rootViper.BindPFlags(cmd.Flags()); err != nil {
		//	return err
		//}
		return do(rootViper)
	},
}

func init() {
	logrus.Trace("doCmd.init")
	defer func() {
		logrus.Trace("doCmd.init ok")
	}()

		rootCmd.AddCommand(doCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	doCmd.Flags().BoolP("doFlag", "d", false, "flag for do")
}

func do(v *viper.Viper) error {
	logrus.Info("all settings:\t", v.AllSettings())
	return fmt.Errorf("some error on do")
}
