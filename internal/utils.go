package internal

import "github.com/fatih/color"

func ColorForType(episodeType EpisodeType) *color.Color {
	switch episodeType {
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
