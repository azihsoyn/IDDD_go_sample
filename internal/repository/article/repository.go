// +build !mock

package articlerepo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/azihsoyn/IDDD_go_sample/internal/domain/article"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/azihsoyn/lwcache.v1"
)

type articleRepository struct {
	db           *sql.DB
	articleCache lwcache.Cache
}

const (
	username = "root"
	password = ""
	host     = "mysql"
	port     = 3306
	dbname   = "iddd_sample_go"
)

var _ article.Repository = (*articleRepository)(nil)

func New() article.Repository {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		username,
		password,
		host,
		port,
		dbname,
	))
	if err != nil {
		panic(err)
	}
	return &articleRepository{
		db: db,
	}

}

func (repo *articleRepository) ResolveByID(articleID article.Identifier) (article.Article, error) {
	as, err := repo.resolveArticleByIDFromMySQL(articleID)
	if err != nil {
		return article.Article{}, err
	}
	return as, nil
}

func (repo *articleRepository) ResolveAll() ([]article.Article, error) {
	return nil, errors.New("not implemented")
}
