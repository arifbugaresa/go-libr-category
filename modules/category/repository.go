package category

import (
	"context"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"go-libr-category/modules/category/pb"
	"go-libr-category/utils/constant"
	"go-libr-category/utils/logger"
	"time"
)

type Repository interface {
	InsertCategory(ctx context.Context, category *pb.Category) (err error)
	UpdateCategory(ctx context.Context, category *pb.Category) (err error)
	ListCategory(ctx context.Context) (res []*pb.Category, err error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &repository{
		db: database,
	}
}

func (r *repository) InsertCategory(ctx context.Context, category *pb.Category) (err error) {
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
func (r *repository) UpdateCategory(ctx context.Context, category *pb.Category) (err error) {
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
			goqu.Ex{"id": category.Id},
		)

	_, err = dataset.Executor().Exec()
	if err != nil {
		logger.ErrorWithCtx(nil, nil, err)
		return
	}

	return
}

func (r *repository) ListCategory(ctx context.Context) (res []*pb.Category, err error) {
	conn := goqu.New(constant.Postgres.Dialect(), r.db)
	dataset := conn.From(constant.Category.TableName()).
		Select(
			goqu.I("id"),
			goqu.I("name"),
			goqu.I("description"),
		)

	rows, err := dataset.Executor().Query()
	if err != nil {
		logger.ErrorWithCtx(nil, nil, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category pb.Category
		if err = rows.Scan(&category.Id, &category.Name, &category.Description); err != nil {
			logger.ErrorWithCtx(nil, nil, err)
			return nil, err
		}
		res = append(res, &category)
	}

	return
}
