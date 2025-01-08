package admin

import (
	"context"
)

type Service interface {
	Login(ctx context.Context, req *LoginRequest) (string, error)
	Register(ctx context.Context, req *LoginRequest) (string, error)
}

type ServiceImpl struct{}

func (s *ServiceImpl) Login(ctx context.Context, req *LoginRequest) (string, error) {
	return "", nil
}

func (s *ServiceImpl) Register(ctx context.Context, req *LoginRequest) (string, error) {
	return "", nil
}
