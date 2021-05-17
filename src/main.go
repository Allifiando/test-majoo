package main

import (
	"log"
	"os"
	"strconv"
	"time"

	_userEntity "test-majoo/src/api/entity/user"
	_userHandler "test-majoo/src/api/handler/user"
	_userRepo "test-majoo/src/api/repo/user"

	"test-majoo/src/config"
	"test-majoo/src/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
			log.Fatal("Error getting env")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "1111"
	}

	timeout := os.Getenv("TIMEOUT")
	if timeout == "" {
		timeout = "2"
	}

	i, _ := strconv.Atoi(timeout)
	timeoutContext := time.Duration(i) * time.Second

	config.Init()
	db := config.GetDB()

	userRepo := _userRepo.InitUserRepo(db)
	userEntity := _userEntity.InitUserEntity(userRepo, timeoutContext)

	r := gin.Default()
	if os.Getenv("ENV") != "local" {
		r.Use(middleware.CORSMiddleware())
	}

	api := r.Group("/")

	_userHandler.InitUserHandler(api, userEntity)

	r.Run(":" + port)
}
