package service

import (
	"context"
	"github.com/list412/tweets-tg-bot/internal/commands"
	"github.com/list412/tweets-tg-bot/internal/config"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	io_prometheus_client "github.com/prometheus/client_model/go"
	"log"
	"time"
)

type Repository interface {
	IsExist(ctx context.Context, userName string) (bool, error)
	Add(ctx context.Context, userName string) error
	Delete(ctx context.Context, userName string) error
	Get(ctx context.Context, userName string) error
	All(ctx context.Context, limit int, offset int) error
	Count(ctx context.Context) (int64, error)
	CountByTime(ctx context.Context, t time.Time) (int, error)
	RefreshDate(ctx context.Context, userName string) error
}

type ShareRepository interface {
	Repository
	RefreshDate(ctx context.Context, userName string) error
	CountByTime(ctx context.Context, t time.Time) (int, error)
}

type MetricHandler interface {
	HandleCmd(ctx context.Context, cmd commands.Cmd)
	GetCmdStats(ctx context.Context, command string) (prometheus.Counter, error)
}

func New(repository Repository, shareRepo ShareRepository, metrics MetricHandler, cfg config.Admin) *Service {
	return &Service{users: repository, usersShare: shareRepo, metrics: metrics, cfg: cfg}
}

type Service struct {
	users      Repository
	usersShare ShareRepository
	metrics    MetricHandler
	cfg        config.Admin
}

func (s Service) Command(cmd commands.Cmd, userName string) {
	ctx := context.TODO()
	switch cmd {
	case commands.InstagramCmd:
		fallthrough
	case commands.TweetCmd:
		fallthrough
	case commands.TikTokCmd:
		err := s.AddShare(ctx, userName)
		if err != nil {
			log.Printf("AddShare error: %s", err.Error())
		}
	default:
		err := s.Add(ctx, userName)
		if err != nil {
			log.Printf("Add error: %s", err.Error())
		}
	}

	s.metrics.HandleCmd(ctx, cmd)
}

func (s Service) IsExist(ctx context.Context, userName string) bool {
	exist, _ := s.users.IsExist(ctx, userName)
	return exist
}

func (s Service) Add(ctx context.Context, userName string) error {
	if !s.IsExist(ctx, userName) {
		err := s.users.Add(ctx, userName)
		if err != nil {
			return errors.Wrap(err, "Add")
		}
	} else {
		err := s.users.RefreshDate(ctx, userName)
		if err != nil {
			return errors.Wrap(err, "refresh date")
		}
	}
	return nil
}

func (s Service) AddShare(ctx context.Context, userName string) error {
	exist, _ := s.usersShare.IsExist(ctx, userName)
	if !exist {
		err := s.usersShare.Add(ctx, userName)
		if err != nil {
			return errors.Wrap(err, "Add")
		}
	} else {
		err := s.usersShare.RefreshDate(ctx, userName)
		if err != nil {
			return errors.Wrap(err, "RefreshDate")
		}
	}
	return nil
}

func (s Service) GetAdminId() int {
	return s.cfg.Id
}

func (s Service) IsAdmin(userId int) (bool, error) {
	if userId == 0 {
		return false, nil
	}
	return userId == s.cfg.Id, nil
}

func (s Service) Count(ctx context.Context) (int, error) {
	count, err := s.users.Count(ctx)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (s Service) CountShare(ctx context.Context) (int, error) {
	count, err := s.usersShare.Count(ctx)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (s Service) CountActiveUsers(ctx context.Context) (int, int, error) {
	m := time.Now().AddDate(0, -1, 0)
	d := time.Now().AddDate(0, 0, -1)

	mau, err := s.usersShare.CountByTime(ctx, m)
	if err != nil {
		return 0, 0, err
	}
	dau, err := s.usersShare.CountByTime(ctx, d)
	if err != nil {
		return mau, 0, err
	}
	return mau, dau, nil
}

func (s Service) CountPassiveUsers(ctx context.Context) (int, int, error) {
	m := time.Now().AddDate(0, -1, 0)
	d := time.Now().AddDate(0, 0, -1)

	mau, err := s.users.CountByTime(ctx, m)
	if err != nil {
		return 0, 0, err
	}
	dau, err := s.users.CountByTime(ctx, d)
	if err != nil {
		return mau, 0, err
	}
	return mau, dau, nil
}

func (s Service) CommandsStat(ctx context.Context) (map[string]int, error) {
	commandsToSend := []string{string(commands.TweetCmd), string(commands.TikTokCmd), string(commands.InstagramCmd), string(commands.HelpCmd), string(commands.StartCmd)}
	result := make(map[string]int, len(commandsToSend))
	for _, c := range commandsToSend {
		counter, err := s.metrics.GetCmdStats(ctx, c)
		if err != nil {
			continue
		}
		store := io_prometheus_client.Metric{}
		err = counter.Write(&store)
		if err != nil {
			continue
		}
		result[c] = int(store.Counter.GetValue())
	}
	return result, nil
}
