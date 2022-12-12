package cmd

import (
	"fmt"

	"github.com/abbit/naruw/internal"
	"github.com/spf13/cobra"
)

func newCurrentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "current",
		Short: "Get information about current episode",
		Long: `This command allows to get information about the current Naruto anime episode you would watch,
including its title, episode number, and whether it is a canon or filler episode.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			episode, err := internal.GetNarutoCurrentEpisode()
			cobra.CheckErr(err)

			fmt.Println("Current Naruto episode:")
			internal.PrintEpisode(&episode)
		},
	}

	return cmd
}
