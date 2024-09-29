package service

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	io_prometheus_client "github.com/prometheus/client_model/go"
	"log"
	"tweets-tg-bot/internal/commands"
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

type MetricHandler interface {
	HandleCmd(ctx context.Context, cmd commands.Cmd)
	GetCmdStats(ctx context.Context, command string) (prometheus.Counter, error)
}

func New(repository Repository, shareRepo Repository, metrics MetricHandler, cfg config.Admin) *Service {
	return &Service{users: repository, usersShare: shareRepo, metrics: metrics, cfg: cfg}
}

type Service struct {
	users      Repository
	usersShare Repository
	metrics    MetricHandler
	cfg        config.Admin
}

func (s Service) Command(cmd commands.Cmd, userName string) {
	ctx := context.TODO()
	switch cmd {
	case commands.TweetCmd:
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
			return err
		}
	}
	return nil
}

func (s Service) AddShare(ctx context.Context, userName string) error {
	exist, _ := s.usersShare.IsExist(ctx, userName)
	if !exist {
		err := s.usersShare.Add(ctx, userName)
		if err != nil {
			return err
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

func (s Service) CommandsStat(ctx context.Context) (map[string]int, error) {
	commandsToSend := []string{string(commands.TweetCmd), string(commands.TikTokCmd), string(commands.HelpCmd), string(commands.StartCmd)}
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
