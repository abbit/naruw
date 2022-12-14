package cmd

import (
	"github.com/abbit/naruw/internal"
	"github.com/spf13/cobra"
)

func newLastCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "last",
		Short: "Get information about last watched episode",
		Long: `This command allows to get information about the last Naruto anime episode you watched,
including its title, episode number, and whether it is a canon or filler episode.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			episode, err := internal.GetNarutoLastWatchedEpisode()
			cobra.CheckErr(err)

			internal.PrintEpisode(&episode)
		},
	}

	return cmd
}
