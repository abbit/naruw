package internal

type EpisodeType string

const (
	MangaCanon       EpisodeType = "Manga Canon"
	AnimeCanon       EpisodeType = "Anime Canon"
	MixedCanonFiller EpisodeType = "Mixed Canon/Filler"
	Filler           EpisodeType = "Filler"
)

type Episode struct {
	Number int         `json:"number"`
	Title  string      `json:"title"`
	Type   EpisodeType `json:"type"`
}
