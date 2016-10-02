// +build !mock

package articlerepo

import (
	"time"

	"github.com/azihsoyn/IDDD_go_sample/internal/domain/article"
)

func (repo *articleRepository) resolveArticleByIDFromMySQL(articleID article.Identifier) (article.Article, error) {
	query := `
	SELECT
		id, 
		title, 
		content
	FROM
		articles
	WHERE 
		id = ?`

	var id int64
	var title, content string
	err := repo.db.QueryRow(query, articleID).Scan(&id, &title, &content)
	if err != nil {
		return article.Article{}, err
	}

	return article.Article{
		ID:      article.Identifier(id),
		Title:   title,
		Content: content,
	}, nil
}

func (repo *articleRepository) resolveArticleByIDFromCache(articleID article.Identifier) (article.Article, bool) {
	if c, ok := repo.articleCache.Get(articleID); ok {
		return c.(article.Article), true
	}

	return article.Article{}, false
}

func (repo *articleRepository) storeArticleToCache(a article.Article) {
	repo.articleCache.Set(a.ID, a, 5*time.Minute)
}
