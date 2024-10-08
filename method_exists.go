package repository

import (
	"errors"

	"github.com/edkirin/gormfilterrepo/smartfilter"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ExistsMethod[T schema.Tabler] struct {
	repo *RepoBase[T]
}

func (m *ExistsMethod[T]) Init(repo *RepoBase[T]) {
	m.repo = repo
}

func (m ExistsMethod[T]) Exists(filter interface{}) (bool, error) {
	var (
		model T
		res   int
	)

	query := m.repo.dbConn.Model(model)

	query, err := smartfilter.ToQuery(model, filter, query)
	if err != nil {
		return false, err
	}

	result := query.Select("1").Take(&res)

	exists := !errors.Is(result.Error, gorm.ErrRecordNotFound) && result.Error == nil
	return exists, nil
}
