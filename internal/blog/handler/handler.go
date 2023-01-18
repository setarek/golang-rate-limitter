package handler

import (
	"golang-rate-limitter/config"
	"golang-rate-limitter/internal/blog/repository"
	"golang-rate-limitter/pkg/logger"
)

type Handler struct {
	Config          *config.Config
	Logger          logger.Logger
	Repository      *repository.BlogRepository
	RedisRepository *repository.RedisRepository
}

func NewHandler(config *config.Config, logger logger.Logger, repository *repository.BlogRepository, redisRepository *repository.RedisRepository) *Handler {
	return &Handler{
		Config:          config,
		Logger:          logger,
		Repository:      repository,
		RedisRepository: redisRepository,
	}
}
