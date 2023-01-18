package handler

import (
	"github.com/gin-gonic/gin"

	"golang-rate-limitter/internal/blog/middleware"
)

func (h *Handler) Register(r *gin.Engine) {
	v1 := r.Group("/api/v1/blog")
	v1.Use(middleware.RateLimitter(h.Logger, h.RedisRepository))
	v1.POST("", h.NewBlog)
}
