package biz

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/weblfe/flyfire/api/account/service/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Account is a Greeter model.
type Account struct {
	ID        string                 `json:"id"`
	Username  string                 `json:"username"`

	CreatedAt *timestamppb.Timestamp `json:"created_at,omitempty"`
	// 开始时间
	UpdatedAt *timestamppb.Timestamp `json:"updated_at,omitempty"`
}

// AccountRepo is an Account repo.
type AccountRepo interface {
	FindByID(context.Context, string) (*Account, error)
}

// AccountUseCase is an account useCase.
type AccountUseCase struct {
	repo AccountRepo
	log  *log.Helper
}

// NewAccountUseCase new a user Account useCase.
func NewAccountUseCase(repo AccountRepo, logger log.Logger) *AccountUseCase {
	return &AccountUseCase{repo: repo, log: log.NewHelper(logger)}
}

// FindByID find a User Account Info, and returns the new Account.
func (uc *AccountUseCase) FindByID(ctx context.Context, id string) (*Account, error) {
	uc.log.WithContext(ctx).Infof("FindByID: %v", id)
	return uc.repo.FindByID(ctx, id)
}
