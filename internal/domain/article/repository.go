package article

//go:generate mockgen -package article -source repository.go -destination mock/repository_mock.go -imports .=github.com/azihsoyn/IDDD_go_sample/internal/domain/article

type Repository interface {
	ResolveByID(articleID Identifier) (Article, error)
	ResolveAll() ([]Article, error)
}
