package cmd

import (
	"github.com/abbit/naruw/internal"
	"github.com/spf13/cobra"
)

func newNextCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "next",
		Short: "Mark current episode as watched and get information about next episode",
		Long: `This command allows to mark the current Naruto episode as watched and get information about next episode,
including its title, episode number, and whether it is a canon or filler episode.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			err := internal.MarkNarutoEpisodeAsWatched()
			cobra.CheckErr(err)

			episode, err := internal.GetNarutoCurrentEpisode()
			cobra.CheckErr(err)

			internal.PrintEpisode(&episode)
		},
	}

	return cmd
}
