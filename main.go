package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

func clean(srv *server, cancel context.CancelFunc) {
	cancel()
}

func run() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	log.SetOutput(&lumberjack.Logger{
		Filename:   home + "/logs/portfolio/access.log",
		MaxSize:    100,
		MaxBackups: 10,
	})
	srv := NewServer()
	go func() {
		if err := srv.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer clean(srv, cancel)
	return srv.Server.Shutdown(ctx)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
