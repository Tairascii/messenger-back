package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"messenger/chats/internal/config"
	httphandler "messenger/chats/internal/http"
	"messenger/chats/internal/http/chats"
	chatsrepo "messenger/chats/internal/repository/chats"
	chatsparticipantsrepo "messenger/chats/internal/repository/chatsparticipants"
	"messenger/chats/internal/usecase/chats/deletechat"
	"messenger/chats/internal/usecase/chats/userchats"
	sharedconfig "messenger/shared/config"
	"messenger/shared/db"
	"messenger/shared/logger"

	"github.com/jmoiron/sqlx"
)

func main() {
	ctx := context.Background()
	cfg, err := sharedconfig.LoadConfig[config.Config]()
	if err != nil {
		logger.Log.Errorf("sharedconfig.LoadConfig: %w", err)
		return
	}

	dbSettings := db.Settings{
		Host:         cfg.DB.Host,
		Port:         cfg.DB.Port,
		User:         cfg.DB.Port,
		Password:     cfg.DB.Password,
		DbName:       cfg.DB.DBName,
		Schema:       cfg.DB.Shema,
		AppName:      cfg.DB.AppName,
		MaxIdleConns: cfg.DB.MaxIdleConns,
		MaxOpenConns: cfg.DB.MaxOpenConns,
	}

	sqlxDb, err := db.Connect(dbSettings)
	if err != nil {
		logger.Log.Errorf("db.Connect: %w", err)
		return
	}
	defer func(sqlxDb *sqlx.DB) {
		if err := sqlxDb.Close(); err != nil {
			logger.Log.Errorf("sqlxDb.Close: %w", err)
		}
	}(sqlxDb)

	chatsRepo := chatsrepo.New(sqlxDb)
	chatsParticipantsRepo := chatsparticipantsrepo.New(sqlxDb)

	userChatsUseCase := userchats.New(&userchats.Config{
		ChatsRepo: chatsRepo,
	})

	deleteChatUseCase := deletechat.New(&deletechat.Config{
		ChatsParticipantsRepo: chatsParticipantsRepo,
	})

	chatsHandlers := chats.New(chats.HandlerConfig{
		UserChatsUseCase:  userChatsUseCase,
		DeleteChatUseCase: deleteChatUseCase,
	})

	handlers := httphandler.New(&httphandler.Config{
		ChatsHandlers: chatsHandlers,
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Service.Port),
		ReadTimeout:  cfg.Service.ReadTimeout,
		WriteTimeout: cfg.Service.WriteTimeout,
		IdleTimeout:  cfg.Service.IdleTimeout,
		Handler:      handlers.InitHandlers(),
	}

	go func() {
		if err := srv.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
			logger.Log.Errorf("srv.ListenAndServe: %w", err)
		}
	}()

	logger.Log.Infof("starting server on port %s", cfg.Service.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-quit

	logger.Log.Info("shutting down server")

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Errorf("srv.Shutdown: %w", err)
	}
}
