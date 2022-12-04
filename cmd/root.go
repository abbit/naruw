package cmd

import (
	"fmt"
	"os"

	"github.com/abbit/narutoep/internal"
	"github.com/abbit/narutoep/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "narutoep",
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

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
}
