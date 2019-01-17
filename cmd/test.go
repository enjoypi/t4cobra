package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// testCmd represents the do command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "this is test command, do not use as sample",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	PreRunE: preRunE,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runSub(rootViper)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	testCmd.Flags().String("test.str", "test string", "string flag for test")
	testCmd.Flags().BoolP("test.bool", "b", false, "bool flag for test")
}

func runSub(v *viper.Viper) error {
	logrus.Info("settings on test: ", v.AllSettings())
	return nil
}
