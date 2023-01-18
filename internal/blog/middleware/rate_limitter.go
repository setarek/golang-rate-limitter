package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"golang-rate-limitter/internal/blog/repository"
	"golang-rate-limitter/pkg/logger"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func RateLimitter(logger logger.Logger, redisRepository *repository.RedisRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		bucketSize, err := redisRepository.GetBucketSize()
		if err != nil {
			logger.Error("Error while getting bucket size from redis: %v", err)
			respondWithError(c, http.StatusInternalServerError, "internal server error")
			return
		}

		if bucketSize == 0 {
			logger.Error("these is too many request to process, try again later")
			respondWithError(c, http.StatusTooManyRequests, "too many request")
			return
		}

		if err := redisRepository.RemoveToken(); err != nil {
			logger.Error("error while dropping one token")
		}

		c.Next()
	}
}
