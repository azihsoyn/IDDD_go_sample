package repository

import (
	"context"
	"testing"

	"github.com/azihsoyn/IDDD_go_sample/internal/domain/article"
	articlemock "github.com/azihsoyn/IDDD_go_sample/internal/domain/article/mock"
	"github.com/golang/mock/gomock"
)

type _contextKey string

var contextKey _contextKey = "repository"

type Repository struct {
	ArticleRepository article.Repository
}

func newRepository(ctx context.Context) *Repository {
	t := new(testing.T)
	ctrl := gomock.NewController(t)
	repo := articlemock.NewMockRepository(ctrl)
	repo.EXPECT().ResolveByID(article.Identifier(1)).Return(article.Article{
		ID:      1,
		Title:   "hello",
		Content: "world",
	}, nil).Times(10)

	repo.EXPECT().ResolveAll().Return([]article.Article{
		{
			ID:      1,
			Title:   "hello",
			Content: "world1",
		},
		{
			ID:      2,
			Title:   "hello",
			Content: "world2",
		},
	}, nil).Times(10)

	return &Repository{
		ArticleRepository: repo,
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
