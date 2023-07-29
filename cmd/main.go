package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/amookia/arvan-backend-challenge/internal/config"
	m "github.com/amookia/arvan-backend-challenge/internal/repository/mongodb"
	"github.com/amookia/arvan-backend-challenge/internal/repository/redis"
	"github.com/amookia/arvan-backend-challenge/internal/service/middleware"
	"github.com/amookia/arvan-backend-challenge/internal/service/upload"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/gin"
	"github.com/amookia/arvan-backend-challenge/pkg/logger/zap"
	"github.com/amookia/arvan-backend-challenge/pkg/mongodb"
	"go.uber.org/zap/zapcore"
)

func main() {
	var conf config.Config
	config.ReadYAML("./config.yaml", &conf)
	file, _ := os.OpenFile("./logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger := zap.New(file, zapcore.ErrorLevel)
	context := context.Background()

	// redis repository
	redisRepo, _ := redis.New(conf.Redis, context, logger)

	// mongodb
	mongo, err := mongodb.ConnectToMongo(
		fmt.Sprintf("mongodb://%s:%s", conf.Mongo.Host, conf.Mongo.Port), conf.Mongo.Name)
	if err != nil {
		log.Fatal(err)
	}
	mongodbRepo := m.New(mongo, logger, context)

	mdService := middleware.New(conf.Middleware, redisRepo, logger)
	upService := upload.New(logger, mongodbRepo)
	serv := gin.New(logger, conf.Middleware, mdService, upService)
	serv.Start(":8080")

}
