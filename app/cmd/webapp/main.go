package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mindera/ready-set-go-workshop/app/internal/webapp"
	"github.com/op/go-logging"
)

const (
	timeout = 30
	port    = 8080
)

var (
	log    = logging.MustGetLogger("webapp")
	stop   = make(chan os.Signal)
	wg     sync.WaitGroup
	client *redis.Client
)

func init() {
	/**
	 ** a) INSERT YOUR CODE BELOW
	 ** Initialise redis client
	 **/
	client = nil
}

func main() {
	r := mux.NewRouter()
	registry := &webapp.RouteRegistry{r, client}

	// register http routes to handlers
	registry.RegisterRoutes()

	recoveryHandler := handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  time.Second * timeout,
		WriteTimeout: time.Second * timeout,
		IdleTimeout:  time.Second * timeout,
		Handler:      recoveryHandler(r),
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Info(fmt.Sprintf("Web application is listening on port %d", port))
		/**
		 ** b) INSERT YOUR CODE BELOW
		 ** http server should listen and serve. check for errors.
		 **/
	}()

	log.Info("Web application is starting")
	<-stop // wait for any signal from OS to stop application
	log.Info("Web application is shutting down...")
	server.Shutdown(context.Background())
	wg.Wait() // wait for any processes to end
	log.Info("Web application gracefully stopped")
}
