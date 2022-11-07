package data

import (
	"context"

	"githu.com/weblfe/flyfire/app/account/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
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
	// r.data.GetDb(ctx).Where()
	return &biz.Account{
			ID:       "",
			Username: "",
	}, nil
}
