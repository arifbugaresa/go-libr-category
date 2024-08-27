package email

import (
	"database/sql"
	"github.com/doug-martin/goqu"
	"github.com/gin-gonic/gin"
	"go-libr-category/utils/constant"
	"go-libr-category/utils/logger"
)

type Repository interface {
	GetEmailTemplate(ctx *gin.Context, code string) (EmailTemplate, error)
}

type emailRepository struct {
	db *sql.DB
}

func NewRepository(dbParam *sql.DB) Repository {
	return &emailRepository{
		db: dbParam,
	}
}

func (r *emailRepository) GetEmailTemplate(ctx *gin.Context, code string) (res EmailTemplate, err error) {
	conn := goqu.New(constant.Postgres.Dialect(), r.db)
	dialect := conn.From(constant.EmailTemplateTableName.TableName()).
		Select(
			goqu.I("id"),
			goqu.I("code"),
			goqu.I("name"),
			goqu.I("template"),
		).
		Where(
			goqu.I("code").Eq(code),
		)

	_, err = dialect.ScanStruct(&res)
	if err != nil {
		logger.ErrorWithCtx(ctx, nil, err)
		return
	}

	return
}
