package cmd

import (
	"strconv"

	"github.com/abbit/naruw/internal"
	"github.com/spf13/cobra"
)

func newInfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info [episode number]",
		Short: "Get information about a specific episode",
		Long: `This command allows to get information about a specific episode of the Naruto anime, 
including its title, episode number, and whether it is a canon or filler episode.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// parse the episode number argument
			episodeNum, err := strconv.Atoi(args[0])
			cobra.CheckErr(err)

			episode, err := internal.GetNarutoEpisode(episodeNum)
			cobra.CheckErr(err)

			internal.PrintEpisode(&episode)
		},
	}

	return cmd
}
