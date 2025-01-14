package user

import (
	"context"
)

type Service interface {
	Login(ctx context.Context, req *LoginRequest) (string, error)
	Register(ctx context.Context, req *LoginRequest) (string, error)
	ListMenu(ctx context.Context, req *[]int) ([]*MenuRequest, error)
}

type ServiceImpl struct{}

func (s *ServiceImpl) Login(ctx context.Context, req *LoginRequest) (string, error) {
	return "", nil
}

func (s *ServiceImpl) Register(ctx context.Context, req *LoginRequest) (string, error) {
	return "", nil
}

func (s *ServiceImpl) ListMenu(ctx context.Context, req *[]int) ([]*MenuRequest, error) {
	return nil, nil
}
