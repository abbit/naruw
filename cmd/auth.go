package cmd

import (
	"github.com/abbit/narutoep/internal/shikimori"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate to Shikimori with OAuth2",
	Run: func(cmd *cobra.Command, args []string) {
		shikimori.Authenticate()
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
