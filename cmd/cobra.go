package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github.com/bulutcan99/weekly-task-scheduler",
	Short: "Developer Scheduler",
	Long: "Developer Scheduler \n" +
		"For http server use " + color.CyanString("`github.com/bulutcan99/weekly-task-scheduler http_server`\n") +
		"For seeder use " + color.CyanString("`github.com/bulutcan99/weekly-task-scheduler seeder`\n"),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
