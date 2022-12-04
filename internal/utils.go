package internal

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/abbit/narutoep/internal/shikimori"
	"github.com/fatih/color"
)

const (
	narutoEpisodesInfoJSONPath = "/Users/abbit/dev/projects/narutoep/data/naruto-p1-episodes-info.json" // TODO: make it generic
	shikimoriNarutoAnimeID     = 41932591
)

type EpisodeType string

const (
	MangaCanon       EpisodeType = "Manga Canon"
	AnimeCanon       EpisodeType = "Anime Canon"
	MixedCanonFiller EpisodeType = "Mixed Canon/Filler"
	Filler           EpisodeType = "Filler"
)

func colorFromEpisodeType(episodeType *EpisodeType) *color.Color {
	switch *episodeType {
	case MangaCanon, AnimeCanon:
		return color.New(color.FgHiGreen)
	case MixedCanonFiller:
		return color.New(color.FgHiYellow)
	case Filler:
		return color.New(color.FgHiRed)
	default:
		return color.New(color.FgWhite)
	}
}

type Episode struct {
	Number  int         `json:"number"`
	Title   string      `json:"title"`
	Type    EpisodeType `json:"type"`
	WikiUrl string      `json:"wikiUrl"`
}

func PrintEpisode(episode *Episode) {
	c := colorFromEpisodeType(&episode.Type)
	c.Printf("Episode #%d (%s)\n", episode.Number, episode.Type)
	fmt.Printf("Title: %s\n", episode.Title)
	fmt.Printf("Wiki link: %s\n", episode.WikiUrl)
}

func getNarutoUserRate() (shikimori.ShikimoriUserRate, error) {
	return shikimori.GetUserRate(shikimoriNarutoAnimeID)
}

func GetNarutoEpisodes() ([]Episode, error) {
	// open the JSON file
	file, err := os.Open(narutoEpisodesInfoJSONPath)
	if err != nil {
		return nil, fmt.Errorf("unable to open JSON file: %w", err)
	}
	defer file.Close()

	// read and decode the JSON data
	var episodes []Episode
	if err := json.NewDecoder(file).Decode(&episodes); err != nil {
		return nil, fmt.Errorf("unable to decode JSON data: %w", err)
	}

	return episodes, nil
}

func GetNarutoEpisode(episodeNum int) (Episode, error) {
	episodes, err := GetNarutoEpisodes()
	if err != nil {
		return Episode{}, err
	}

	// search for the episode with the given number
	var episode Episode
	for _, ep := range episodes {
		if ep.Number == episodeNum {
			episode = ep
			break
		}
	}

	if episode.Number == 0 {
		return Episode{}, fmt.Errorf("episode not found")
	}

	return episode, nil
}

func GetNarutoLastWatchedEpisode() (Episode, error) {
	userRate, err := getNarutoUserRate()
	if err != nil {
		return Episode{}, err
	}

	return GetNarutoEpisode(userRate.Episodes)
}

func GetNarutoCurrentEpisode() (Episode, error) {
	userRate, err := getNarutoUserRate()
	if err != nil {
		return Episode{}, err
	}

	return GetNarutoEpisode(userRate.Episodes + 1)
}

func MarkNarutoEpisodeAsWatched() error {
	return shikimori.IncrementEpisodes(shikimoriNarutoAnimeID)
}
