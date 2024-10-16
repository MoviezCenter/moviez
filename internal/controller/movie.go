package controller

import (
	"context"

	"github.com/MoviezCenter/pb-contracts-go/core"
	moviepb "github.com/MoviezCenter/pb-contracts-go/movie"
)

type MovieServiceServer struct {
	moviepb.UnimplementedMovieServiceServer
}

func NewMovieServiceServer() *MovieServiceServer {
	return &MovieServiceServer{}
}

func (s *MovieServiceServer) GetMovieList(context.Context, *moviepb.GetMovieListRequest) (*moviepb.GetMovieListResponse, error) {
	return &moviepb.GetMovieListResponse{
		Data: []*core.Movie{},
	}, nil
}
