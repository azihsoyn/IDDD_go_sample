package article

import (
	"net/http"

	"github.com/azihsoyn/IDDD_go_sample/internal/domain/article"
)

type articleView struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func renderArticleView(w http.ResponseWriter, a article.Article) {
	view := articleView{
		ID:      int64(a.ID),
		Title:   a.Title,
		Content: a.Content,
	}
	renderer.JSON(w, http.StatusOK, view)
}
