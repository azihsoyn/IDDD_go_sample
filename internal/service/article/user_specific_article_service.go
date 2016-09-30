package article

import "github.com/azihsoyn/IDDD_go_sample/internal/domain/article"

type userSpecificArticleService struct {
	repo article.Repository
}

func NewUserSpecificArticleService(articleRepo article.Repository) article.UserSpecificArticleService {
	return &userSpecificArticleService{
		repo: articleRepo,
	}
}

func (s *userSpecificArticleService) Get(userID int64) ([]article.Article, error) {
	as, err := s.repo.ResolveAll()
	if err != nil {
		return nil, err
	}

	ret := make([]article.Article, 0, len(as))
	if userID == 1 {
		ret = append(ret, article.Article{
			ID:      100,
			Title:   "for user 1",
			Content: "this is recommended by your activity",
		})
	}
	ret = append(ret, as...)

	return ret, nil
}
