package service

import (
	"context"

	"github.com/MoviezCenter/moviez/internal/domain/model"
	"github.com/MoviezCenter/moviez/internal/repository"
	"github.com/MoviezCenter/pb-contracts-go/core"
)

type MovieService interface {
	GetMovies(ctx context.Context, limit int, offset int) ([]*core.Movie, error)
}

type movieService struct {
	movieRepo repository.MovieRepository
}

func NewMovieService(movieRepo repository.MovieRepository) MovieService {
	return &movieService{
		movieRepo: movieRepo,
	}
}

func (s *movieService) GetMovies(ctx context.Context, limit int, offset int) ([]*core.Movie, error) {
	movies, err := s.movieRepo.GetMovies(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var res []*core.Movie
	for _, m := range movies {
		res = append(res, model.ToProtoMovie(m))
	}

	return res, nil
}
