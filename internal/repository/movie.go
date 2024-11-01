package repository

import (
	"context"

	"github.com/MoviezCenter/moviez/ent"
)

type MovieRepository interface {
	GetMovies(ctx context.Context, limit int, offset int) ([]*ent.Movie, error)
}

type movieRepo struct {
	db *ent.Client
}

func NewMovieRepo(db *ent.Client) MovieRepository {
	return &movieRepo{
		db: db,
	}
}

func (r *movieRepo) GetMovies(ctx context.Context, limit int, offset int) ([]*ent.Movie, error) {
	if limit <= 0 {
		limit = 10
	}
	movies, err := r.db.Movie.Query().
		Limit(limit).
		Offset(offset).
		WithGenres().
		All(ctx)
	if err != nil {
		return nil, err
	}

	return movies, nil
}
