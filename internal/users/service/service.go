package service

import (
	"context"
	"tweets-tg-bot/internal/config"
)

type Repository interface {
	IsExist(ctx context.Context, userName string) (bool, error)
	Add(ctx context.Context, userName string) error
	Delete(ctx context.Context, userName string) error
	Get(ctx context.Context, userName string) error
	All(ctx context.Context, limit int, offset int) error
	Count(ctx context.Context) (int64, error)
}

func New(repository Repository, shareRepo Repository, cfg config.Admin) *service {
	return &service{users: repository, usersShare: shareRepo, cfg: cfg}
}

type service struct {
	users      Repository
	usersShare Repository
	cfg        config.Admin
}

func (s service) IsExist(ctx context.Context, userName string) bool {
	exist, _ := s.users.IsExist(ctx, userName)
	return exist
}

func (s service) Add(ctx context.Context, userName string) error {
	if !s.IsExist(ctx, userName) {
		err := s.users.Add(ctx, userName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s service) AddShare(ctx context.Context, userName string) error {
	exist, _ := s.usersShare.IsExist(ctx, userName)
	if !exist {
		err := s.usersShare.Add(ctx, userName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s service) IsAdmin(userId int) (bool, error) {
	if userId == 0 {
		return false, nil
	}
	return userId == s.cfg.Id, nil
}

func (s service) Count(ctx context.Context) (int, error) {
	count, err := s.users.Count(ctx)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (s service) CountShare(ctx context.Context) (int, error) {
	count, err := s.usersShare.Count(ctx)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
