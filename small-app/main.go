package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"small-app/handlers"
	"small-app/internal/users"
	"time"
)

// run go mod tidy // to install all the dependencies
// Goal: User would send a userEmail and we would look up the user and return the user's details'

/*
Graceful Shutdown:

	Once requested, no new requests will be accepted.
	Existing requests will be completed, but a certain amount of time will be allowed for them to complete.
*/
func main() {
	// initializing the map
	// we should not call this function inside the handler functions
	// handler functions can be called millions of times
	// we want to initialize our dependencies only once, not every time we call of handler function
	con := users.NewConn()

	// overriding the default http server, with timeout values and other configurations
	api := http.Server{
		Addr:         ":" + "3000",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		IdleTimeout:  800 * time.Second,
		// handler.InitRoutes returns a *gin.Engine, which implements the http.Handler interface
		Handler: handlers.InitRoutes(con),
	}

	serverError := make(chan error)
	// starting the server in a goroutine, so it doesn't block the main'
	go func() {
		err := api.ListenAndServe()
		serverError <- err
	}()

	shutdown := make(chan os.Signal, 1)

	//singal.Notify registers the given channel to receive notifications of the specified signals.
	signal.Notify(shutdown, os.Interrupt, os.Kill)

	// using select to wait for either a server error or a shutdown signal
	// if a shutdown signal is received, we will gracefully shutdown the server
	// select would block until either a server error or a shutdown signal is received
	select {
	case err := <-serverError:
		log.Println(err)
		return
	case <-shutdown:
		log.Println("Gracefully Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		//Shutdown gracefully shuts down the server without interrupting any active connections.
		//Shutdown works by first closing all open listeners, then closing all idle connections,
		err := api.Shutdown(ctx)
		if err != nil {

			// forceful shutdown
			err := api.Close()
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
		}
		return
	}

}
