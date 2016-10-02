// +build mock

package articlerepo

import (
	"testing"

	"github.com/azihsoyn/IDDD_go_sample/internal/domain/article"
	articlemock "github.com/azihsoyn/IDDD_go_sample/internal/domain/article/mock"
	"github.com/golang/mock/gomock"
)

func New() article.Repository {
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

	return repo
}
