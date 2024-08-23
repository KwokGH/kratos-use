package cache

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type ICache interface {
	Set(ctx context.Context, key string, data interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, out interface{}) error
	Del(ctx context.Context, key string) error

	HSet(ctx context.Context, region string, cond interface{}, data interface{}, expiration time.Duration) error
	HGet(ctx context.Context, region string, cond interface{}, out interface{}) error
	HDelField(ctx context.Context, region string, cond interface{}) error
	HDelKey(ctx context.Context, region string) error

	GetKey(ctx context.Context, cond interface{}) string
}

type Config struct {
	Protocol    Protocol
	Address     string
	ProjectName string
	AttachTls   bool
}

type Protocol string

const (
	ProtocolRedis Protocol = "redis"
)

func CreateClient(cfg Config) (ICache, error) {
	switch cfg.Protocol {
	default:
		return NewRedisCache(cfg)
	}
}

func FieldKey(condition interface{}) string {
	h := md5.New()
	b, _ := json.Marshal(condition)
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

var (
	ErrNotFound = errors.New("redis key not found")
)
