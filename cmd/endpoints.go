package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(traderCmd)
}

var traderCmd = &cobra.Command{
	Use:   "trader",
	Short: "Starting candle receiver",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
