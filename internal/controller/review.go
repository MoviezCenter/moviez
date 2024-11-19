package controller

import (
	"context"

	"github.com/MoviezCenter/moviez/internal/service"
	"github.com/MoviezCenter/moviez/pkg/response"
	"github.com/MoviezCenter/pb-contracts-go/core"
	reviewpb "github.com/MoviezCenter/pb-contracts-go/review"
)

type ReviewServiceServer struct {
	reviewSvc service.ReviewService
	reviewpb.UnimplementedReviewServiceServer
}

func NewReviewServiceServer(reviewSvc service.ReviewService) *ReviewServiceServer {
	return &ReviewServiceServer{
		reviewSvc: reviewSvc,
	}
}

func (c *ReviewServiceServer) GetReviewList(ctx context.Context, req *reviewpb.GetReviewListRequest) (*reviewpb.GetReviewListResponse, error) {
	reviews, total, err := c.reviewSvc.ListReviews(ctx, int(req.MovieId), int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	return &reviewpb.GetReviewListResponse{
		Total: uint32(total),
		Data:  reviews,
	}, nil
}

func (c *ReviewServiceServer) AddReview(ctx context.Context, req *reviewpb.AddReviewRequest) (*core.GeneralResponse, error) {
	_, err := c.reviewSvc.CreateReview(ctx, int(req.MovieId), 1, float32(req.Rating), req.Review)
	if err != nil {
		return nil, err
	}

	return &core.GeneralResponse{
		Code:    response.ResponseSuccessCode,
		Message: "Created review successfully",
	}, nil
}

func (c *ReviewServiceServer) UpdateReview(ctx context.Context, req *reviewpb.UpdateReviewRequest) (*core.GeneralResponse, error) {
	err := c.reviewSvc.UpdateReview(ctx, int(req.ReviewId), int(req.MovieId), 1, float32(req.Rating), req.Review)
	if err != nil {
		return nil, err
	}
	return &core.GeneralResponse{
		Code:    response.ResponseSuccessCode,
		Message: "Updated review successfully",
	}, nil
}

func (c *ReviewServiceServer) DeleteReview(ctx context.Context, req *reviewpb.DeleteReviewRequest) (*core.GeneralResponse, error) {
	err := c.reviewSvc.DeleteReview(ctx, int(req.ReviewId), int(req.MovieId), 1)
	if err != nil {
		return nil, err
	}
	return &core.GeneralResponse{
		Code:    response.ResponseSuccessCode,
		Message: "Deleted review successfully",
	}, nil
}
