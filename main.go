package main

import (
	"fmt"

	"golang-rate-limitter/config"
	"golang-rate-limitter/internal/blog/cronjob"
	"golang-rate-limitter/internal/blog/handler"
	"golang-rate-limitter/internal/blog/repository"
	"golang-rate-limitter/internal/blog/router"
	"golang-rate-limitter/pkg/logger"
	"golang-rate-limitter/pkg/postgres"
	"golang-rate-limitter/pkg/redis"
)

func init() {
	config, err := config.InitConfig()
	if err != nil {
		panic(fmt.Errorf("error while initializing config: %v+", err))
	}
	redisClient := redis.GetRedisClient(config)
	redisRepository := repository.NewRedisRepository(redisClient)
	redisRepository.RefillBucket(config.GetInt64("bucket_size"), 60)
}

func main() {
	config, err := config.InitConfig()
	if err != nil {
		panic(fmt.Errorf("error while initializing config: %v+", err))
	}

	logger := logger.NewLogger(config)
	logger.InitLogger()

	db, err := postgres.InitDB(config, logger)
	if err != nil {
		panic(fmt.Errorf("error while initializing mongo: %v+", err))
	}

	redisClient := redis.GetRedisClient(config)
	redisRepository := repository.NewRedisRepository(redisClient)
	repository := repository.NewRepository(db)

	handler := handler.NewHandler(config, logger, repository, redisRepository)

	r := router.New()
	handler.Register(r)

	go cronjob.InitializeCronJobs(config.GetInt64("refill_rate"), config.GetInt64("interval"), logger, redisRepository)
	defer cronjob.StopCronJobs()

	r.Run(fmt.Sprintf("%s:%s", config.GetString("hostname"), config.GetString("port")))

}
