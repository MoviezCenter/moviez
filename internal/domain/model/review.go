package model

import (
	"github.com/MoviezCenter/moviez/ent"
	"github.com/MoviezCenter/pb-contracts-go/core"
)

func ToProtoReview(review *ent.Review) *core.Review {
  return &core.Review{
    Id: uint32(review.ID),
    Review: review.Comment,
    Rating: uint32(review.Rating),
  }
}
