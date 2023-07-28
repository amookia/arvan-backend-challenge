package main

import (
	"context"
	"log"
	"os"

	"github.com/amookia/arvan-backend-challenge/internal/config"
	"github.com/amookia/arvan-backend-challenge/internal/repository/redis"
	"github.com/amookia/arvan-backend-challenge/internal/service/middleware"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/gin"
)

func main() {
	var conf config.Config
	config.ReadYAML("./config.yaml", &conf)
	file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	context := context.Background()
	// redis repository
	redis, _ := redis.New(conf.Redis, context, logger)
	mdService := middleware.New(conf.Middleware, redis, logger)
	serv := gin.New(logger, conf.Middleware, mdService)
	serv.Start(":8085")

}
