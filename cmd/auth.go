package cmd

import (
	"github.com/abbit/naruw/internal/shikimori"
	"github.com/spf13/cobra"
)

func newAuthCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate to Shikimori with OAuth2",
		Run: func(cmd *cobra.Command, args []string) {
			shikimori.Authenticate()
		},
	}

	return cmd
}
