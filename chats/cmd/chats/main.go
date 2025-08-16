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

	"github.com/jmoiron/sqlx"
)

func main() {
	ctx := context.Background()
	cfg, err := sharedconfig.LoadConfig[config.Config]()
	if err != nil {
		// log
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
		// log
		return
	}
	defer func(sqlxDb *sqlx.DB) {
		if err := sqlxDb.Close(); err != nil {
			// log
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
			// log
		}
	}()

	// log

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-quit

	// log

	if err := srv.Shutdown(ctx); err != nil {
		// err
	}
}
