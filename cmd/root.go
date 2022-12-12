package cmd

import (
	"fmt"
	"os"

	"github.com/abbit/naruw/internal"
	"github.com/abbit/naruw/internal/config"
	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "naruw",
		Short: "CLI for managing Naruto anime episodes",
		Long: `This CLI app allows to get information about Naruto episodes, 
including whether the episode is a canon or filler episode,
get last watched episode and mark episodes as watched on Shikimori.

If called without any arguments, it will print information about the current episode.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			episode, err := internal.GetNarutoCurrentEpisode()
			cobra.CheckErr(err)

			fmt.Println("Current Naruto episode:")
			internal.PrintEpisode(&episode)
		},
	}

	cmd.AddCommand(newAuthCommand())
	cmd.AddCommand(newInfoCommand())
	cmd.AddCommand(newLastCommand())
	cmd.AddCommand(newNextCommand())

	return cmd
}

func Execute() {
	// initialize config before executing the command
	cobra.OnInitialize(config.InitConfig)

	if err := newRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
