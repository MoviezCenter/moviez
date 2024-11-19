package service

import (
	"context"
	"errors"

	"github.com/MoviezCenter/moviez/ent"
	"github.com/MoviezCenter/moviez/internal/domain/model"
	"github.com/MoviezCenter/moviez/internal/repository"
	"github.com/MoviezCenter/pb-contracts-go/core"
)

type ReviewService interface {
	ListReviews(ctx context.Context, movieID, limit, offset int) ([]*core.Review, int, error)
	CreateReview(ctx context.Context, movieID int, creatorID int, rating float32, comment string) (int, error)
	UpdateReview(ctx context.Context, reviewID int, movieID int, creatorID int, rating float32, comment string) error
	DeleteReview(ctx context.Context, reviewID int, movieID int, creatorID int) error
}

type reviewService struct {
	reviewRepo repository.ReviewRepository
}

func NewReviewService(reviewRepo repository.ReviewRepository) ReviewService {
	return &reviewService{
		reviewRepo: reviewRepo,
	}
}

func (s *reviewService) ListReviews(ctx context.Context, movieID, limit, offset int) ([]*core.Review, int, error) {
	reviews, err := s.reviewRepo.ListReviews(ctx, movieID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.reviewRepo.GetTotalReviews(ctx, movieID)
	if err != nil {
		return nil, 0, err
	}

	var result []*core.Review
	for _, r := range reviews {
		result = append(result, model.ToProtoReview(r))
	}

	return result, total, nil
}

func (s *reviewService) CreateReview(ctx context.Context, movieID int, creatorID int, rating float32, comment string) (int, error) {
	review, err := s.reviewRepo.GetReviewByUserID(ctx, movieID, creatorID)
	if err != nil && !ent.IsNotFound(err) {
		return 0, err
	}

  if review != nil {
    return 0, errors.New("error user has already reviewed this movie")
  }

	reviewID, err := s.reviewRepo.CreateReview(ctx, movieID, creatorID, rating, comment)
	if err != nil {
		return 0, err
	}

	return reviewID, nil
}

func (s *reviewService) UpdateReview(ctx context.Context, reviewID int, movieID int, creatorID int, rating float32, comment string) error {
	err := s.reviewRepo.UpdateReview(ctx, reviewID, movieID, creatorID, rating, comment)
	if err != nil {
		return err
	}

	return nil
}

func (s *reviewService) DeleteReview(ctx context.Context, reviewID int, movieID int, creatorID int) error {
	err := s.reviewRepo.DeleteReview(ctx, reviewID, movieID, creatorID)
	if err != nil {
		return err
	}

	return nil
}
