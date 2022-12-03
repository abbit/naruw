package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "narutoep",
	Short: "CLI for managing Naruto anime episodes",
	Long: `This CLI app allows users to get information about Naruto episodes, 
including whether the episode is a canon or filler episode. It also allows users to mark 
episodes as watched on a specified website via API.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
