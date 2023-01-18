package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	rsErr "golang-rate-limitter/internal/blog/error"
	"golang-rate-limitter/internal/blog/model"
)

type NewBlogRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (h *Handler) NewBlog(ctx *gin.Context) {

	var request NewBlogRequest

	if err := ctx.Bind(&request); err != nil {
		h.Logger.Error("error while binding body request", err)
		ctx.JSON(http.StatusUnprocessableEntity, ErrorResponse{
			Message: rsErr.ErrEmptyBodyRequest.Error(),
		})

	}

	newBlog := model.BlogModel{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := h.Repository.CreatePost(newBlog); err != nil {
		h.Logger.Error("error while create blog post", err)
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "new blog post created",
	})

}
