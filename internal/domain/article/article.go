package article

type Identifier int64

type Article struct {
	ID      Identifier
	Title   string
	Content string
}

type Articles []Article

func UserSpecificArticle(as []Article, userID int64) []Article {
	ret := make([]Article, 0, len(as))
	if userID == 1 {
		ret = append(ret, Article{
			ID:      100,
			Title:   "for user 1",
			Content: "this is recommended by your activity",
		})
	}
	ret = append(ret, as...)

	return ret
}
