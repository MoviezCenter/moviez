package model

import (
	"github.com/MoviezCenter/moviez/ent"
	"github.com/MoviezCenter/pb-contracts-go/core"
)

func ToProtoMovie(m *ent.Movie) *core.Movie {
	coreMovie := &core.Movie{
		Id:          uint32(m.ID),
		Title:       m.Title,
		Overview:    m.Overview,
		ReleaseDate: m.ReleaseDate,
		PosterPath:  m.PosterPath,
	}

	for _, g := range m.Edges.Genres {
		coreMovie.Genres = append(coreMovie.Genres, ToProtoGenre(g))
	}

	return coreMovie
}

func ToProtoGenre(g *ent.Genre) *core.Genre {
	return &core.Genre{
		Id:   uint32(g.ID),
		Name: g.Name,
		Type: ToProtoType(g.TypeID),
	}
}

func ToProtoType(typeId int32) core.Type {
	switch typeId {
	case int32(core.Type_TYPE_MOVIE):
		return core.Type_TYPE_MOVIE
	case int32(core.Type_TYPE_TV_SHOW):
		return core.Type_TYPE_TV_SHOW
	default:
		return core.Type_TYPE_MOVIE_UNSPECIFIED
	}
}
