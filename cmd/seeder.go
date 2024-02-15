package cmd

import (
	"github.com/bulutcan99/weekly-task-scheduler/cmd/seeder"
	"github.com/spf13/cobra"
)

var seederCmd = &cobra.Command{
	Use:     "seeder",
	Short:   "Seeder",
	Long:    "Seeder",
	Aliases: []string{"seeder", "seeder"},
	Run: func(cmd *cobra.Command, args []string) {
		seeder.Start()
	},
}

func init() {
	rootCmd.AddCommand(seederCmd)
}
