package article

import (
	"net/http"

	"github.com/azihsoyn/IDDD_go_sample/internal/domain/article"
)

type articleListView []articleView

func renderArticleListView(w http.ResponseWriter, as []article.Article) {
	view := make(articleListView, 0, len(as))
	for _, a := range as {
		view = append(view, articleView{
			ID:      int64(a.ID),
			Title:   a.Title,
			Content: a.Content,
		})
	}
	renderer.JSON(w, http.StatusOK, view)
}
