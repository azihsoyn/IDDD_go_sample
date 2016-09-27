package article

import (
	"context"
	"net/http"

	"github.com/azihsoyn/IDDD_go_sample/internal/repository"
)

func ArticleHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := articleRequest{}
	if err := req.Parse(ctx); err != nil {
		renderer.JSON(w, http.StatusBadRequest, err.Error())
	}

	a, err := repository.FromContext(ctx).ArticleRepository.ResolveByID(req.ArticleID)
	if err != nil {
		renderer.JSON(w, http.StatusBadRequest, err.Error())
	}

	renderArticleView(w, a)
}
