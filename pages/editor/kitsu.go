package editor

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/kitsu"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// NewKitsuAnime ...
func NewKitsuAnime(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	finder := arn.NewAnimeFinder("kitsu/anime")

	animes := arn.FilterKitsuAnime(func(anime *kitsu.Anime) bool {
		return finder.GetAnime(anime.ID) == nil
	})

	sort.Slice(animes, func(i, j int) bool {
		a := animes[i]
		b := animes[j]

		return a.ID > b.ID
	})

	return ctx.HTML(components.NewKitsuAnime(animes, ctx.URI(), user))
}
