package article

import (
	"context"
	"net/http"

	"github.com/azihsoyn/IDDD_go_sample/internal/service"
)

func ArticleListHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := articleListRequest{}
	if err := req.Parse(r); err != nil {
		renderer.JSON(w, http.StatusBadRequest, err.Error())
	}

	as, err := service.FromContext(ctx).UserSpecificArticleService.Get(req.UserID)
	if err != nil {
		renderer.JSON(w, http.StatusBadRequest, err.Error())
	}

	renderArticleListView(w, as)
}
