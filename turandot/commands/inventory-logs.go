package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	inventoryCommand.AddCommand(inventoryLogsCommand)
	inventoryLogsCommand.Flags().StringVarP(&component, "component", "c", "registry", "sub-component (\"registry\" or \"spooler\")")
	inventoryLogsCommand.Flags().IntVarP(&tail, "tail", "t", -1, "number of most recent lines to print (<0 means all lines)")
	inventoryLogsCommand.Flags().BoolVarP(&follow, "follow", "f", false, "keep printing incoming logs")
}

var inventoryLogsCommand = &cobra.Command{
	Use:   "logs",
	Short: "Show the logs of the Turandot inventory",
	Run: func(cmd *cobra.Command, args []string) {
		Logs("inventory", component)
	},
}
