package controller

import (
	"go-backend/domain"
	"go-backend/setup"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



type threadController struct {
	threadUsecase domain.ThreadUsecase
	postUsecase domain.PostUsecase
	env *setup.Env
}

func NewThreadController(tu domain.ThreadUsecase, pu domain.PostUsecase, env *setup.Env) *threadController {
	return &threadController{threadUsecase: tu, postUsecase: pu, env: env}
}

func (tc *threadController) FetchOneThread(c *gin.Context) {
	// Validate Request Format
	tid := c.Param("tid")
	page := c.Param("page")
	pageInt64, err := strconv.ParseUint(page, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &domain.ErrorResponse{Message: "Unknown url"})
	}
	tidInt64, err := strconv.ParseUint(tid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &domain.ErrorResponse{Message: "Unknown url"})
	}

	// Fetch thread and posts of current page
	thread, err := tc.threadUsecase.GetThreadByID(c, uint(tidInt64))
	if err != nil {
		if err.Error() == "Not Found" {
			c.JSON(http.StatusNotFound, &domain.ErrorResponse{Message: "Thread not exsit"})
		} else {
			c.JSON(http.StatusInternalServerError, &domain.ErrorResponse{Message: "Server Error: Fetch Thread"})
		}
	}

	posts, err := tc.postUsecase.GetPostsByThreadAndPage(c, uint(tidInt64), uint(pageInt64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &domain.ErrorResponse{Message: "Server Error: Fetch Post"})
	}

	// Update Thread viewCount
	err = tc.threadUsecase.UpdateViewCount(c, uint(tidInt64), 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &domain.ErrorResponse{Message: "Server Error: Update View Count"})
	}

	// Send Response
	response := &domain.FetchOneThreadResponse{Thread: *thread, Posts: posts}
	c.JSON(http.StatusOK, &response)
}

func (tc *threadController) FetchThreads(c *gin.Context) {
	// Validate Request Format
	page := c.Param("page")
	order := c.Param("order")
	pageUint64, err := strconv.ParseUint(page, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &domain.ErrorResponse{Message: "Unknown url"})
	}

	// Fetch Threads
	threads, err := tc.threadUsecase.GetThreadsByPage(c, uint(pageUint64), order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &domain.ErrorResponse{Message: "Server Error: Fetch threads"})
	}

	// Send Response
	response := &domain.FetchThreadsResponse{Threads: threads}
	c.JSON(http.StatusOK, &response)
}

func (tc *threadController) CreateThread(c *gin.Context) {
	// Validate Request Format
	var request domain.CreateThreadRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, &domain.ErrorResponse{Message: "Bad Request"})
	}

	// Write to database
	err = tc.threadUsecase.CreateThread(c, &request.Thread)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &domain.ErrorResponse{Message: "Server Error: Create Thread"})
	}

	err = tc.postUsecase.CreatePost(c, &request.Post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &domain.ErrorResponse{Message: "Server Error: Create Post"})
	}

	// Send Response
	c.JSON(http.StatusOK, nil) 
}

func (tc *threadController) CreatePost(c *gin.Context) {
	// Validate Request Format
	var request domain.CreatePostRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, &domain.ErrorResponse{Message: "Bad Request"})
	}

	// Write to cache and database asyncly
	err = tc.postUsecase.CreatePost(c, &request.Post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &domain.ErrorResponse{Message: "Server Error: Create Post"})
	}

	// Send Response
	c.JSON(http.StatusOK, nil)
}