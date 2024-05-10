package repo

import (
	"github.com/bloodblue999/umhelp/config"
	"github.com/bloodblue999/umhelp/repo/mysql"
	"github.com/bloodblue999/umhelp/repo/redis"
)

type RepoManager struct {
	MySQL *mysql.Repo
	Redis *redis.Repo
}

func New(cfg *config.Config) (*RepoManager, error) {
	mysql, err := mysql.New(cfg)
	if err != nil {
		return nil, err
	}

	redis, err := redis.New(cfg)
	if err != nil {
		return nil, err
	}

	return &RepoManager{
		MySQL: mysql,
		Redis: redis,
	}, nil
}
