package article

type UserSpecificArticleService interface {
	Get(userID int64) ([]Article, error)
}
