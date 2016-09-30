package service

import (
	"context"

	"github.com/azihsoyn/IDDD_go_sample/internal/domain/article"
	"github.com/azihsoyn/IDDD_go_sample/internal/repository"
	articleservice "github.com/azihsoyn/IDDD_go_sample/internal/service/article"
)

type _contextKey string

var contextKey _contextKey = "service"

type Service struct {
	UserSpecificArticleService article.UserSpecificArticleService
}

func newService(ctx context.Context) *Service {
	articleRepo := repository.FromContext(ctx).ArticleRepository
	return &Service{
		UserSpecificArticleService: articleservice.NewUserSpecificArticleService(articleRepo),
	}
}

func NewContext(ctx context.Context) context.Context {
	return WithContext(ctx, newService(ctx))
}

func WithContext(ctx context.Context, srv *Service) context.Context {
	return context.WithValue(ctx, contextKey, srv)
}

func FromContext(ctx context.Context) *Service {
	r, ok := ctx.Value(contextKey).(*Service)
	if !ok {
		r = newService(ctx)
	}
	return r
}
