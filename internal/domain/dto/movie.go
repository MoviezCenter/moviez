package dto

import (
	"github.com/MoviezCenter/moviez/internal/domain/model"
	"github.com/MoviezCenter/pb-contracts-go/core"
)

type MovieDto struct {
	model.Movie
	model.Genre
}

func (m MovieDto) ToProto() *core.Movie {
	return &core.Movie{
		Id:          m.Movie.Id,
		Title:       m.Title,
		Overview:    m.Overview,
		ReleaseDate: m.Releasedate,
		PosterPath:  m.PosterPath,
		Genre: &core.Genre{
			Id:   m.Genre.Id,
			Name: m.Genre.Name,
			Type: ToProtoType(m.Genre.TypeId),
		},
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
