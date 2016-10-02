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

	as = article.UserSpecificArticle(as, userID)
	return as, nil
}
