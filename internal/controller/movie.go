package controller

import (
	"context"

	"github.com/MoviezCenter/moviez/internal/service"
	moviepb "github.com/MoviezCenter/pb-contracts-go/movie"
)

type MovieServiceServer struct {
	movieServie service.MovieService
	moviepb.UnimplementedMovieServiceServer
}

func NewMovieServiceServer(movieServie service.MovieService) *MovieServiceServer {
	return &MovieServiceServer{
		movieServie: movieServie,
	}
}

func (s *MovieServiceServer) GetMovieList(ctx context.Context, req *moviepb.GetMovieListRequest) (*moviepb.GetMovieListResponse, error) {
	movies, err := s.movieServie.GetMovies(ctx, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}
	return &moviepb.GetMovieListResponse{
		Data: movies,
	}, nil
}

func (s *MovieServiceServer) GetMovieDetail(ctx context.Context, req *moviepb.GetMovieDetailRequest) (*moviepb.GetMovieDetailResponse, error) {
	movie, err := s.movieServie.GetMovieDetail(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &moviepb.GetMovieDetailResponse{
		Data: movie,
	}, nil
}
