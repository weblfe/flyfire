package data

import (
		"context"
		"xorm.io/builder"

		"github.com/go-kratos/kratos/v2/log"
		"github.com/weblfe/flyfire/app/account/service/internal/biz"
)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

// NewAccountRepo .
func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *accountRepo) FindByID(ctx context.Context, id string) (*biz.Account, error) {
	r.data.GetDb(ctx).Get(builder.In(""))

	return &biz.Account{
		ID:       "",
		Username: "",
	}, nil
}
