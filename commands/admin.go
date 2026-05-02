package commands

import "github.com/spf13/cobra"

var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "",
}

func init() {
	rootCmd.AddCommand(adminCmd)
}
