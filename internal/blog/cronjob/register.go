package cronjob

import (
	"fmt"

	cron "github.com/robfig/cron/v3"

	"golang-rate-limitter/internal/blog/repository"
	"golang-rate-limitter/pkg/logger"
)

var cronClient *cron.Cron

func init() {
	cronClient = cron.New()
}

func InitializeCronJobs(bucketSize int64, interval int64, logger logger.Logger, repository *repository.RedisRepository) {
	cronClient.AddFunc(fmt.Sprintf("@every %ds", interval), RefillBucket(bucketSize, interval, logger, repository))
	cronClient.Start()
}

func StopCronJobs() {
	cronClient.Stop()
}
