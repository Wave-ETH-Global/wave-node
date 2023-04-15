package middlewares

import (
	"context"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/Wave-ETH-Global/wave-node/controllers/domain"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type SessionStore struct {
	client *redis.Client
}

func NewSessionStore(client *redis.Client) *SessionStore {
	return &SessionStore{client: client}
}

func ProvideRedis(cfg *config.Config) *SessionStore {
	cli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%d",
			cfg.Redis.Host,
			cfg.Redis.Port,
		),
		DB:       int(cfg.Redis.Database),
		Password: "",
	})

	return NewSessionStore(cli)
}

func (s *SessionStore) GetWalletAddress(
	ctx context.Context,
	key string,
) (domain.Address, error) {
	value, err := s.client.Get(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return "", errors.WithStack(err)
	}

	if value == "" {
		return "", nil
	}

	return domain.Address(value), nil
}

const keyLength = 64

func (s *SessionStore) SetWalletAddress(
	ctx context.Context,
	address domain.Address,
) (string, error) {
	key, err := secureRandom(keyLength)
	if err != nil {
		return "", errors.WithStack(err)
	}

	if err := s.client.Set(ctx, key, address, time.Duration(time.Duration.Hours(24*7))).Err(); err != nil {
		return "", errors.WithStack(err)
	}

	return key, nil
}

// nolint:gochecknoglobals
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func secureRandom(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	result := make([]rune, n)
	for i, v := range b {
		result[i] = letterRunes[int(v)%len(letterRunes)]
	}
	return string(result), nil
}
