package category

import (
	"context"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"go-libr-category/utils/constant"
	"go-libr-category/utils/logger"
	"time"
)

type Repository interface {
	InsertCategory(ctx context.Context, category DTOCategory) (err error)
	UpdateCategory(ctx context.Context, category DTOCategory) (err error)
	ListCategory(ctx context.Context) (res []DTOCategory, err error)
	GetCategoryById(ctx context.Context, id int64) (res DTOCategory, err error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &repository{
		db: database,
	}
}

func (r *repository) InsertCategory(ctx context.Context, category DTOCategory) (err error) {
	conn := goqu.New(constant.Postgres.Dialect(), r.db)
	dataset := conn.Insert(constant.Category.TableName()).Rows(
		goqu.Record{
			"name":        category.Name,
			"description": category.Description,
		},
	)

	_, err = dataset.Executor().Exec()
	if err != nil {
		logger.ErrorWithCtx(nil, nil, err)
		return
	}

	return
}
func (r *repository) UpdateCategory(ctx context.Context, category DTOCategory) (err error) {
	conn := goqu.New(constant.Postgres.Dialect(), r.db)

	dataset := conn.Update(constant.Category.TableName()).
		Set(
			goqu.Record{
				"name":        category.Name,
				"description": category.Description,
				"modified_at": time.Now(),
			},
		).
		Where(
			goqu.Ex{"id": category.ID},
		)

	_, err = dataset.Executor().Exec()
	if err != nil {
		logger.ErrorWithCtx(nil, nil, err)
		return
	}

	return
}

func (r *repository) ListCategory(ctx context.Context) (res []DTOCategory, err error) {
	conn := goqu.New(constant.Postgres.Dialect(), r.db)
	dataset := conn.From(constant.Category.TableName()).
		Select(
			goqu.I("id"),
			goqu.I("name"),
			goqu.I("description"),
		)

	err = dataset.ScanStructs(&res)
	if err != nil {
		logger.ErrorWithCtx(nil, nil, err)
		return
	}

	return
}

func (r *repository) GetCategoryById(_ context.Context, id int64) (res DTOCategory, err error) {
	conn := goqu.New(constant.Postgres.Dialect(), r.db)
	dataset := conn.From(constant.Category.TableName()).
		Select(
			goqu.I("id"),
			goqu.I("name"),
			goqu.I("description"),
		).
		Where(
			goqu.Ex{"id": id},
		)

	_, err = dataset.ScanStruct(&res)
	if err != nil {
		logger.ErrorWithCtx(nil, nil, err)
		return
	}

	return res, nil
}
