package article

import (
	"context"
	"strconv"

	"github.com/azihsoyn/IDDD_go_sample/internal/domain/article"
	"github.com/guregu/kami"
)

type articleRequest struct {
	ArticleID article.Identifier
}

func (req *articleRequest) Parse(ctx context.Context) error {
	articleID, err := strconv.ParseInt(kami.Param(ctx, "article_id"), 10, 64)
	if err != nil {
		return err
	}
	req.ArticleID = article.Identifier(articleID)
	return nil
}
