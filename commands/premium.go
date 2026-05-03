package commands

import "github.com/spf13/cobra"

var premiumCmd = &cobra.Command{
	Use:   "premium",
	Short: "",
}

func init() {
	rootCmd.AddCommand(premiumCmd)
}
