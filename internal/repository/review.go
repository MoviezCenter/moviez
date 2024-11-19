package repository

import (
	"context"
	"time"

	"github.com/MoviezCenter/moviez/ent"
	"github.com/MoviezCenter/moviez/ent/review"
)

type ReviewRepository interface {
	ListReviews(ctx context.Context, movieID, limit, offset int) ([]*ent.Review, error)
	GetTotalReviews(ctx context.Context, movieID int) (int, error)
	CreateReview(ctx context.Context, movieID int, creatorID int, rating float32, comment string) (int, error)
	UpdateReview(ctx context.Context, reviewID int, movieID int, creatorID int, rating float32, comment string) error
	DeleteReview(ctx context.Context, reviewID int, movieID int, creatorID int) error
	GetReviewByUserID(ctx context.Context, movieID int, creatorID int) (*ent.Review, error)
}

type reviewRespotory struct {
	db *ent.Client
}

func NewReviewRepository(db *ent.Client) ReviewRepository {
	return &reviewRespotory{
		db: db,
	}
}

func (r *reviewRespotory) ListReviews(ctx context.Context, movieID, limit, offset int) ([]*ent.Review, error) {
	if limit <= 0 {
		limit = 10
	}

	if offset < 0 {
		offset = 0
	}

	reviews, err := r.db.Review.Query().
		Where(review.MovieID(movieID)).
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *reviewRespotory) GetTotalReviews(ctx context.Context, movieID int) (int, error) {
	count, err := r.db.Review.Query().
		Where(review.MovieID(movieID)).
		Count(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *reviewRespotory) CreateReview(ctx context.Context, movieID int, creatorID int, rating float32, comment string) (int, error) {
	review, err := r.db.Review.Create().
		SetMovieID(movieID).
		SetCreatorID(creatorID).
		SetRating(rating).
		SetComment(comment).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).Save(ctx)
	if err != nil {
		return 0, err
	}

	return review.ID, nil
}

func (r *reviewRespotory) UpdateReview(ctx context.Context, reviewID int, movieID int, creatorID int, rating float32, comment string) error {
	_, err := r.db.Review.UpdateOneID(reviewID).
		SetRating(rating).
		SetComment(comment).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *reviewRespotory) DeleteReview(ctx context.Context, reviewID int, movieID int, creatorID int) error {
	err := r.db.Review.DeleteOneID(reviewID).
		Where(review.And(
			review.MovieID(movieID),
			review.CreatorID(creatorID),
		)).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *reviewRespotory) GetReviewByUserID(ctx context.Context, movieID int, creatorID int) (*ent.Review, error) {
	review, err := r.db.Review.Query().
		Where(review.And(
			review.MovieID(movieID),
			review.CreatorID(creatorID),
		)).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return review, nil
}
