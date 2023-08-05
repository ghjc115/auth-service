package main

import (
	authService "auth-service/cmd/auth-service"
	"auth-service/dao"
	"auth-service/internal/config"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// setup logger
	logger := setupLogger()

	logger.Info(fmt.Sprintf("Starting auth-service[%s-%s]", cfg.Version, cfg.Env))

	// connect to mongo-db
	mongoUri := fmt.Sprintf("%s:%s", cfg.Database.Host, cfg.Database.Port)
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to MongoDB. %s", err))

		return
	}

	userDAO, err := dao.NewUserDAO(mongoClient, logger)
	if err != nil {
		logger.Error(fmt.Sprintf("Error in NewUserDAO ctor. %s", err))
	}

	// create authService
	service := authService.NewService(userDAO)

	// create router
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// create httpHandler
	httpHandler := authService.NewHandler(service, logger)

	// init router endpoints
	router.Post("/users/add", httpHandler.Add)
	router.Get("/users/get/{nickname}", httpHandler.Get)
	router.Delete("/users/delete/{nickname}", httpHandler.Delete)

	server := &http.Server{
		Addr:         cfg.HttpServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HttpServer.Timeout,
		WriteTimeout: cfg.HttpServer.Timeout,
		IdleTimeout:  cfg.HttpServer.IdleTimeout,
	}

	logger.Info(fmt.Sprintf("Server started at %s", cfg.Address))

	err = server.ListenAndServe()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to start server. %s", err))

		return
	}
}

func setupLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
}
