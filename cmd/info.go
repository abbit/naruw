package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	. "github.com/abbit/narutoep/internal"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info [episode number]",
	Short: "Get information about a Naruto anime episode",
	Long: `This command allows users to get information about a specific episode of the Naruto anime, 
including its title, episode number, and whether it is a canon or filler episode.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// parse the episode number argument
		episodeNum, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: invalid episode number")
			return
		}

		// open the JSON file
		file, err := os.Open("/Users/abbit/dev/projects/narutoep/data/naruto-p1-episodes-info.json")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: unable to open JSON file")
			return
		}
		defer file.Close()

		// read and decode the JSON data
		var episodes []Episode
		if err := json.NewDecoder(file).Decode(&episodes); err != nil {
			fmt.Fprintln(os.Stderr, "Error: unable to decode JSON data")
			return
		}

		// search for the episode with the given number
		var episode Episode
		for _, ep := range episodes {
			if ep.Number == episodeNum {
				episode = ep
				break
			}
		}

		// print the episode information
		if episode.Number > 0 {
			c := ColorForType(episode.Type)
			c.Printf("Episode #%d (%s)\n%s\n", episode.Number, episode.Type, episode.Title)
		} else {
			fmt.Fprintln(os.Stderr, "Error: episode not found")
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
