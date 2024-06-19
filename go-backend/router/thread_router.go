package router

import (
	"context"
	"encoding/json"
	"fmt"
	"go-backend/controller"
	"go-backend/domain"
	"go-backend/repository"
	"go-backend/setup"
	"go-backend/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func ThreadRouterSetup(env *setup.Env, db domain.Database, rdb *redis.Client, group *gin.RouterGroup) {
	tr := repository.NewThreadRepository(db, rdb, env)
	pr := repository.NewPostRepository(db, rdb, env)
	tu := usecase.NewThreadUsecase(tr, env)
	pu := usecase.NewPostUsecase(pr, env)
	tc := controller.NewThreadController(tu, pu, env)

	go updateHotThreads(tu, pu, rdb, env)

	group.Handle("GET", "/thread/:tid/:page", tc.FetchOneThread)
	group.Handle("GET", "/threads/:page/:order", tc.FetchThreads)
	group.Handle("POST", "/createThread", tc.CreateThread)
	group.Handle("POST", "/createPost", tc.CreatePost)
}

func updateHotThreads(tu domain.ThreadUsecase, pu domain.PostUsecase, rdb *redis.Client, env *setup.Env) {
	// Clear old cache first
	var cursor uint64 = 0
	for {
		keys, newCursor, err := rdb.Scan(context.Background(), cursor, "page*", 10).Result()
		if err != nil {
			log.Println("Error: Fail to clear old cache")
			return
		}
		
		err = rdb.Del(context.Background(), keys...).Err()
		if err != nil {
			log.Println("Error: Fail to clear old cache")
			return
		}
		cursor = newCursor
		if cursor == 0 {
			break
		}
	}

	// Fetch first page of hot thread
	threads, err := tu.GetThreadsByPage(context.Background(), 1, "hot")
	if err != nil {
		log.Println("Error: Fail to fetch hot threads")
		return
	}

	for i := range threads {
		tid := threads[i].ID
		key := fmt.Sprintf("page:%d", tid)
		rdb.JSONSet(context.Background(), key, "$", "[]")  // Empty JSON list
		postCount := threads[i].PostCount
		pageCount := postCount/env.PostsPerPage + 1
		for page := 1; page <= pageCount; page++ {
			posts, err := pu.GetPostsByThreadAndPage(context.Background(), tid, uint(page))
			if err != nil {
				log.Printf("Error: Fail to fetch posts (tid: %d, page: %d)\n", tid, page)
				return
			}

			cacheResponse := &domain.FetchOneThreadResponse{Thread: threads[i], Posts: posts}
			cacheResponseJSON, err := json.Marshal(cacheResponse)
			if err != nil {
				log.Println("Error: Fail to marshal to json")
			}

			// Store cached posts in redis
			err = rdb.JSONArrAppend(context.Background(), key, "$", cacheResponseJSON).Err()
			if err != nil {
				log.Println("Error: Fail to store cache")
				return
			}
		}
	}

}
