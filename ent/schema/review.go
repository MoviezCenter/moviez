package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.String("comment"),
		field.Float32("rating").Max(5.0).Min(0.5),
		field.Int("creator_id").Optional(),
		field.Int("movie_id").Optional(),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", User.Type).Ref("reviews").Field("creator_id").Unique(),
		edge.From("movie", Movie.Type).Ref("reviews").Field("movie_id").Unique(),
	}
}
