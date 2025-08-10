package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	sharedconfig "messenger/shared/config"
	"messenger/shared/db"
	"messenger/user/internal/config"
	httphandler "messenger/user/internal/http"
	"messenger/user/internal/http/auth"
	"messenger/user/internal/repository/user"
	"messenger/user/internal/usecase/auth/signin"
	"messenger/user/internal/usecase/auth/signup"

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

	userRepo := user.New(sqlxDb)

	signInUseCase := signin.New(&signin.Config{
		UserRepo: userRepo,
	})

	signUpUseCase := signup.New(&signup.Config{
		UserRepo: userRepo,
	})

	authHandlers := auth.New(auth.HandlerConfig{
		SignInUseCase: signInUseCase,
		SignUpUseCase: signUpUseCase,
	})

	handlers := httphandler.New(&httphandler.Config{
		AuthHandlers: authHandlers,
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
