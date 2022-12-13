package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abbit/naruw/internal/shikimori"
	"github.com/fatih/color"
)

const (
	narutoEpisodesInfoJsonURL = "https://raw.githubusercontent.com/abbit/naruw/main/data/naruto-p1-episodes-info.json"
	shikimoriNarutoAnimeID    = 41932591
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
	// get episodes info from URL
	response, err := http.Get(narutoEpisodesInfoJsonURL)
	if err != nil {
		return nil, fmt.Errorf("unable to get episodes info: %w", err)
	}
	defer response.Body.Close()

	// read and decode episodes info
	var episodes []Episode
	if err := json.NewDecoder(response.Body).Decode(&episodes); err != nil {
		return nil, fmt.Errorf("unable to decode episodes info: %w", err)
	}

	// TODO: add caching

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
