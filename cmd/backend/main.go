package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/igor-karpukhin/user-api-test/pkg/config"
	"github.com/igor-karpukhin/user-api-test/pkg/server"
	"github.com/igor-karpukhin/user-api-test/pkg/user"
	"go.uber.org/zap"
)

func main() {
	appConfig := config.FromFlags()
	appCtx, cancelF := context.WithCancel(context.Background())

	l, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("unable to init logger: %s\r\n", err.Error())
		os.Exit(1)
	}

	l.Info("application config", zap.Any("config", appConfig))

	mkUserDao := user.NewMockedUserDao()
	mkUserDao.SetUsers(map[uint64]*user.User{
		1: &user.User{
			ID:       1,
			Name:     "Test1",
			Birthday: time.Now(),
			Age:      0,
			Hobbies:  []string{"Sleep"},
		},
	})
	userService := user.NewUserService(mkUserDao, l.Named("user_service"))

	hRouter := mux.NewRouter()
	userService.GetRoutes(hRouter)

	hServer := server.NewHTTPServer(appConfig.HttpHost, appConfig.HttpPort, hRouter, l.Named("http_server"))

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGTERM)

	hServer.Start(appCtx)
	l.Info("server started")

	for {
		select {
		case <-sigCh:
			cancelF()
		}
	}
}
