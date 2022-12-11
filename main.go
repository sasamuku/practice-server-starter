package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/lestrrat/go-server-starter/listener"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)

	listeners, err := listener.ListenAll()
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Handler: http.HandlerFunc(displayPid),
	}

	go server.Serve(listeners[0])

	<-signals
	server.Shutdown(context.Background())
}

func displayPid(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "server pid: %d\n", os.Getpid())
}
