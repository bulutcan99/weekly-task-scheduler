package cmd

import (
	"github.com/bulutcan99/weekly-task-scheduler/cmd/http_server"
	"github.com/spf13/cobra"
)

var httpServerCmd = &cobra.Command{
	Use:     "http_server",
	Short:   "Fiber HTTP Server",
	Long:    "Fiber HTTP Server",
	Aliases: []string{"http_server", "httpServer"},
	Run: func(cmd *cobra.Command, args []string) {
		http_server.Start()
	},
}

func init() {
	rootCmd.AddCommand(httpServerCmd)
}
