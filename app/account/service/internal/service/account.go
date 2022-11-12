package service

import (
	"context"

	"github.com/weblfe/flyfire/api/account/service/v1"
	"github.com/weblfe/flyfire/app/account/service/internal/biz"
)

// AccountService is an account service.
type AccountService struct {
	v1.UnimplementedAccountServer

	uc *biz.AccountUseCase
}

// NewAccountService new an account service.
func NewAccountService(uc *biz.AccountUseCase) *AccountService {
	return &AccountService{uc: uc}
}

// GetUserInfo implements account.GetUserInfo.
func (s *AccountService) GetUserInfo(ctx context.Context, in *v1.GetUserInfoParams) (*v1.GetUserInfoReply, error) {
	g, err := s.uc.FindByID(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetUserInfoReply{
		Id:        g.ID,
		Username:  g.Username,
		RoleType:  v1.RoleType_Normal_USER,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}, nil
}
