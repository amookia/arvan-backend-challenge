package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/amookia/arvan-backend-challenge/internal/config"
	mongo "github.com/amookia/arvan-backend-challenge/internal/repository/mongodb"
	"github.com/amookia/arvan-backend-challenge/internal/repository/redis"
	"github.com/amookia/arvan-backend-challenge/internal/service/middleware"
	"github.com/amookia/arvan-backend-challenge/internal/service/upload"
	"github.com/amookia/arvan-backend-challenge/internal/transport/http/gin"
	"github.com/amookia/arvan-backend-challenge/pkg/logger/zap"
	"github.com/amookia/arvan-backend-challenge/pkg/mongodb"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Logger
	file, _ := os.OpenFile("./logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	logger := zap.New(file, zapcore.ErrorLevel)
	//Config
	var conf config.Config
	config.ReadYAML("./config.yaml", &conf)
	err := config.ReadEnv(&conf)
	if err != nil {
		panic(err)
	}
	// Context
	context := context.Background()
	// Redis
	redisRepo, _ := redis.New(conf.Redis, context, logger)

	// Mongo
	mongoConn, err := mongodb.ConnectToMongo(
		fmt.Sprintf("mongodb://%s:%s", conf.Mongo.Host, conf.Mongo.Port),
		conf.Mongo.Name)
	if err != nil {
		log.Fatal(err)
	}
	mongodbRepo := mongo.New(mongoConn, logger, context)

	// Services
	mdService := middleware.New(conf.Middleware, redisRepo, logger)
	upService := upload.New(logger, mongodbRepo, redisRepo)
	serv := gin.New(logger, conf.Middleware, mdService, upService)
	// Listen
	listenPort := fmt.Sprintf(":%s",conf.Service.Port)
	err = serv.Start(listenPort)
	if err != nil {
		panic(err)
	}
	logger.Info("listen","server start listening at port ",listenPort)

}