package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/amookia/arvan-backend-challenge/internal/config"
	"github.com/amookia/arvan-backend-challenge/internal/repository/redis"
	"github.com/amookia/arvan-backend-challenge/internal/service/middleware"
)

func main() {
	var conf config.Config
	config.ReadYAML("./config.yaml", &conf)
	file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	ctx := context.Background()
	redis, _ := redis.New(conf.Redis, ctx, logger)
	md := middleware.New(conf.Middleware, redis, logger)
	for {
		i := md.IsUserLimited("user")
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
