package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

//1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
func startHttpServer() *http.Server {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello world")
	})

	s := &http.Server{
		Addr:           ":8082",
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}

func listenSignal() os.Signal {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	// Stop the service gracefully.
	return <-ch
}

func main() {
	s := startHttpServer()
	var shutDown func() error
	g, ctx := errgroup.WithContext(context.Background())
	shutDown = func() error {
		listenSignal()
		fmt.Println("\n shut down")
		return s.Shutdown(ctx)
	}
	//http服务
	g.Go(s.ListenAndServe)
	//接受信号 退出服务
	g.Go(shutDown)
	// Wait gorotine print shutdown message
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully shutdown two servers")
	}

}
