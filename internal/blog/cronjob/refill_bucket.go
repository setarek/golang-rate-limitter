package cronjob

import (
	"golang-rate-limitter/internal/blog/repository"
	"golang-rate-limitter/pkg/logger"
)

func RefillBucket(bucketSize int64, interval int64, logger logger.Logger, redisRepository *repository.RedisRepository) func() {
	return func() {
		err := redisRepository.RefillBucket(bucketSize, interval)
		if err != nil {
			logger.Error("Error while rifilling bucket %v", err)
		}
		logger.Info("Bucket just refilled")
	}
}
