package repository

import (
	"context"

	"github.com/azihsoyn/IDDD_go_sample/internal/domain/article"
	articlerepo "github.com/azihsoyn/IDDD_go_sample/internal/repository/article"
)

type _contextKey string

var contextKey _contextKey = "repository"

type Repository struct {
	ArticleRepository article.Repository
}

func newRepository(ctx context.Context) *Repository {
	return &Repository{
		ArticleRepository: articlerepo.New(),
	}
}

func NewContext(ctx context.Context) context.Context {
	return WithContext(ctx, newRepository(ctx))
}

func WithContext(ctx context.Context, repo *Repository) context.Context {
	return context.WithValue(ctx, contextKey, repo)
}

func FromContext(ctx context.Context) *Repository {
	r, ok := ctx.Value(contextKey).(*Repository)
	if !ok {
		r = newRepository(ctx)
	}
	return r
}
