package main

import (
	"context"

	"github.com/azihsoyn/IDDD_go_sample/article"
	"github.com/azihsoyn/IDDD_go_sample/internal/repository"
	"github.com/guregu/kami"
)

func main() {
	ctx := context.Background()
	ctx = repository.NewContext(ctx)
	kami.Context = ctx
	kami.Get("/articles", article.ArticleListHandler)
	kami.Get("/article/:article_id", article.ArticleHandler)
	kami.Serve()
}
