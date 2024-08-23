package cache

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCache struct {
	cfg    Config
	client *redis.Client
}

func NewRedisCache(cfg Config) (*RedisCache, error) {
	rdsOption := &redis.Options{
		Addr:         cfg.Address,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	if cfg.AttachTls {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		rdsOption.TLSConfig = tlsConfig
	}

	rdb := redis.NewClient(rdsOption)
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Infof("redis 连接失败, %s", err.Error())
		return nil, err
	} else {
		log.Info("redis 连接成功")
	}

	return &RedisCache{
		cfg:    cfg,
		client: rdb,
	}, nil
}

func (r *RedisCache) Set(ctx context.Context, key string, data interface{}, expiration time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	key = r.GetKey(ctx, key)
	err = r.client.Set(ctx, key, string(b), expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCache) Get(ctx context.Context, key string, out interface{}) error {
	key = r.GetKey(ctx, key)

	res, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		log.Context(ctx).Warnw("msg", "redis 缓存未命中", "key", key)
		return ErrNotFound
	}
	if err != nil {
		log.Context(ctx).Errorw("msg", "redis hget error", "key", key)
		return err
	}

	return json.Unmarshal([]byte(res), out)
}

func (r *RedisCache) Del(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCache) HSet(ctx context.Context, region string, cond interface{}, data interface{}, expiration time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	field := r.GetKey(ctx, cond)
	err = r.client.HSet(ctx, region, field, string(b)).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCache) HGet(ctx context.Context, region string, cond interface{}, out interface{}) error {
	field := r.GetKey(ctx, cond)

	res, err := r.client.HGet(ctx, region, field).Result()
	if errors.Is(err, redis.Nil) {
		log.Context(ctx).Warnw("msg", "redis 缓存未命中", "region", region, "field", field)
		return ErrNotFound
	}
	if err != nil {
		log.Context(ctx).Errorw("msg", "redis hget error", "region", region, "field", field)
		return err
	}

	return json.Unmarshal([]byte(res), out)
}

func (r *RedisCache) HDelKey(ctx context.Context, region string) error {
	err := r.client.Del(ctx, region).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCache) HDelField(ctx context.Context, region string, cond interface{}) error {
	key := r.GetKey(ctx, cond)

	err := r.client.HDel(ctx, region, key).Err()
	if err != nil {
		return err
	}

	return nil
}
func (r *RedisCache) GetKey(ctx context.Context, cond interface{}) string {
	if cond == nil {
		return ""
	}

	if str, ok := cond.(string); ok {
		return r.cfg.ProjectName + str
	} else {
		return r.cfg.ProjectName + FieldKey(cond)
	}
}
