package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"log"

	"github.com/toshim45/demo-gqlgen/graph/model"
)

// Thing is the resolver for the thing field.
func (r *queryResolver) Thing(ctx context.Context, limit *int, offset *int, where *model.ThingsBoolExp) ([]*model.Thing, error) {
	zero := 0
	if limit == nil {
		limit = &zero
	}
	if offset == nil {
		offset = &zero
	}
	log.Println("[query-Thing] limit-offset", *limit, *offset)
	if where != nil {
		if where.Name != nil {
			if where.Name.Eq != nil {
				log.Println("[query-Thing] bool exp Name eq", *where.Name.Eq)
			}
		}
	}
	return []*model.Thing{
		{ID: "01917255-aabd-7908-a5e0-f4912fc98f03"},
		{ID: "01917256-e6f2-76ca-8c08-db07e9ff72e8"}}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
