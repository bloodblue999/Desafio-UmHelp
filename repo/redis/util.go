package redis

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type Util struct {
	cli *redis.Client
}

func (u *Util) SetString(ctx context.Context, k string, v string, e time.Duration) error {
	if _, err := u.cli.Set(ctx, k, v, e).Result(); err != nil {
		return err
	}

	return nil
}

func (u *Util) SetStructure(ctx context.Context, k string, v interface{}, e time.Duration) error {
	vBytes, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return u.SetString(ctx, k, string(vBytes), e)
}

func (u *Util) GetString(ctx context.Context, k string) (string, error) {
	v, err := u.cli.Get(ctx, k).Result()
	if err != nil {
		if err == redis.Nil {
			return "", errors.New("key not exists in REDIS")
		}

		return "", err
	}

	return v, nil
}

func (u *Util) DelString(ctx context.Context, k string) error {
	_, err := u.cli.Del(ctx, k).Result()
	return err
}
