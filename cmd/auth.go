package cmd

import (
	"github.com/abbit/naruw/internal/shikimori"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate to Shikimori with OAuth2",
	Run: func(cmd *cobra.Command, args []string) {
		/* TODO: почему-то логин не работает больше 1 дня
		истекает срок токена?
		прочитать доки shikmori и ouath2 */
		shikimori.Authenticate()
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
