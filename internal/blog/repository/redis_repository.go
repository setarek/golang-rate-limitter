package repository

import (
	"time"

	"github.com/go-redis/redis"

	"golang-rate-limitter/pkg/utils"
)

type RedisRepository struct {
	rc *redis.Client
}

func NewRedisRepository(rc *redis.Client) *RedisRepository {
	return &RedisRepository{rc: rc}
}

func (r *RedisRepository) RefillBucket(bucketSize interface{}, expiration int64) error {
	x := 60 * time.Second
	if err := r.rc.Set("bucket:size", bucketSize, x).Err(); err != nil {
		return err
	}
	return nil
}

func (r *RedisRepository) GetBucketSize() (int64, error) {
	bucketSize, err := r.rc.Get("bucket:size").Result()
	if err != nil {
		return 0, err
	}
	return utils.ParseInt64(bucketSize), nil
}

func (r *RedisRepository) RemoveToken() error {
	if _, err := r.rc.Decr("bucket:size").Result(); err != nil {
		return err
	}
	return nil
}
