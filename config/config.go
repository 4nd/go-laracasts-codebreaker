package config

import (
	"code-breaker/views"
	"context"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

func Init() {
	loadEnv()

	err := views.LoadViews()
	if err != nil {
		log.Fatal(err)
	}
}

func SetupRedisSessions() *redisstore.RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_SERVER") + ":" + os.Getenv("REDIS_PORT"),
	})

	store, err := redisstore.NewRedisStore(context.Background(), client)
	if err != nil {
		log.Fatal("Failed to create redis store: ", err)
	}

	store.KeyPrefix("session_")
	store.Options(sessions.Options{
		Path:   "/",
		Domain: "localhost",
		MaxAge: 86400,
	})

	return store
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %s", err)
	}
}
