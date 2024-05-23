package main

import (
	_ "database/sql"
	"go-backend/database"
	"go-backend/domain"
	"go-backend/middleware"
	"go-backend/router"
	"go-backend/setup"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	/* Load Environment Variable */
	env := setup.NewEnv()

	/* Connect to database */
	passwd := env.DBpassword
	dsn := "host=localhost user=postgres dbname=forumdb_test port=5432 sslmode=disable password=" + passwd
	db, err := database.NewDatabase(dsn)
	if err != nil {
		log.Fatal(err)
		return
	}

	if env.TestMode {
		tx := db.Begin()
		defer tx.Rollback()
		db = database.NewDatabaseFromExist(tx)
	}

	/* Run database migration if set in env */
	if env.RunMigration {
		db.AutoMigrate(&domain.User{})
	}
	
	/* Setup router */
	ginEngine := gin.Default()

	corsMiddleware := cors.New(cors.Config{
        AllowAllOrigins: true,
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
    })
	ginEngine.Use(corsMiddleware)

	publicRouter := ginEngine.Group("")
	privateRouter := ginEngine.Group("")
	jwtMiddleware := middleware.NewJWTmiddleware(env)
	privateRouter.Use(jwtMiddleware.GinHandler)

	router.SignupRouterSetup(env, db, publicRouter)
	router.LoginRouterSetup(env, db, publicRouter)
	router.ProfileRouterSetup(env, db, privateRouter)

	/* Run */
	ginEngine.Run(":8080")
}