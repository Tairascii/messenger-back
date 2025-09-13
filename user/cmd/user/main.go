package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	sharedconfig "messenger/shared/config"
	"messenger/shared/db"
	"messenger/shared/logger"
	"messenger/user"
	"messenger/user/internal/config"
	httphandler "messenger/user/internal/http"
	"messenger/user/internal/http/auth"
	"messenger/user/internal/http/profile"
	userrepo "messenger/user/internal/repository/user"
	"messenger/user/internal/usecase/auth/signin"
	"messenger/user/internal/usecase/auth/signup"
	"messenger/user/internal/usecase/profile/userprofile"

	"github.com/flowchartsman/swaggerui"
	"github.com/jmoiron/sqlx"
)

//	@title			User
//	@version		1.0.0
//	@description	User service

//	@host		localhost:8080
//	@BasePath	/

func main() {
	ctx := context.Background()
	cfg, err := sharedconfig.LoadConfig[config.Config]()
	if err != nil {
		logger.Log.Errorf("sharedconfig.LoadConfig: %s", err.Error())
		return
	}

	dbSettings := db.Settings{
		Host:         cfg.DB.Host,
		Port:         cfg.DB.Port,
		User:         cfg.DB.User,
		Password:     cfg.DB.Password,
		DbName:       cfg.DB.DBName,
		Schema:       cfg.DB.Shema,
		AppName:      cfg.DB.AppName,
		MaxIdleConns: cfg.DB.MaxIdleConns,
		MaxOpenConns: cfg.DB.MaxOpenConns,
	}

	sqlxDb, err := db.Connect(dbSettings)
	if err != nil {
		logger.Log.Errorf("db.Connect: %s", err.Error())
		return
	}
	defer func(sqlxDb *sqlx.DB) {
		if err := sqlxDb.Close(); err != nil {
			logger.Log.Errorf("sqlxDb.Close: %s", err.Error())
		}
	}(sqlxDb)

	userRepo := userrepo.New(sqlxDb)

	signInUseCase := signin.New(&signin.Config{
		UserRepo: userRepo,
	})

	signUpUseCase := signup.New(&signup.Config{
		UserRepo: userRepo,
	})

	userProfileUseCase := userprofile.New(&userprofile.Config{
		UserRepo: userRepo,
	})

	authHandlers := auth.New(auth.HandlerConfig{
		SignInUseCase: signInUseCase,
		SignUpUseCase: signUpUseCase,
	})

	profileHandlers := profile.New(profile.HandlerConfig{
		UserProfileUseCase: userProfileUseCase,
	})

	handlers := httphandler.New(&httphandler.Config{
		AuthHandlers:    authHandlers,
		ProfileHandlers: profileHandlers,
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Service.Port),
		ReadTimeout:  cfg.Service.ReadTimeout,
		WriteTimeout: cfg.Service.WriteTimeout,
		IdleTimeout:  cfg.Service.IdleTimeout,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/docs") {
				http.StripPrefix("/docs", swaggerui.Handler(user.Spec)).ServeHTTP(w, r)
				return
			}
			handlers.InitHandlers().ServeHTTP(w, r)
		}),
	}

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Errorf("srv.ListenAndServe: %s", err.Error())
		}
	}()

	logger.Log.Infof("starting server on port %s", cfg.Service.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-quit

	logger.Log.Info("shutting down server")

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Errorf("srv.Shutdown: %s", err.Error())
	}
}
