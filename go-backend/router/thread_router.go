package router

import (
	"go-backend/controller"
	"go-backend/domain"
	"go-backend/repository"
	"go-backend/setup"
	"go-backend/usecase"

	"github.com/gin-gonic/gin"
)

func ThreadRouterSetup(env *setup.Env, db domain.Database, group *gin.RouterGroup) {
	tr := repository.NewThreadRepository(db, env)
	pr := repository.NewPostRepository(db, env)
	tu := usecase.NewThreadUsecase(tr, env)
	pu := usecase.NewPostUsecase(pr, env)
	tc := controller.NewThreadController(tu, pu, env)

	group.Handle("GET", "/thread/:tid/:page", tc.FetchOneThread)
	group.Handle("GET", "/threads/:page/:order", tc.FetchThreads)
	group.Handle("POST", "/createThread", tc.CreateThread)
	group.Handle("POST", "/createPost", tc.CreatePost)
}