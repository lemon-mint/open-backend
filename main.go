package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lemon-mint/open-backend/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/valyala/fasthttp"
)

var infolog = log.New(os.Stderr, "[INFO]: ", log.Default().Flags())
var errlog = log.New(os.Stderr, "[ERROR]: ", log.Default().Flags())

func main() {
	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	initDB(client)

	// Create a new server
	server := &fasthttp.Server{
		ReadTimeout:     time.Second * 15,
		WriteTimeout:    time.Minute * 15,
		IdleTimeout:     time.Second * 15,
		CloseOnShutdown: true,
	}

	r := routes()
	server.Handler = r.Handler

	go func() {
		if os.Getenv("PORT") != "" {
			infolog.Println("Server starting on port", os.Getenv("PORT"))
			err := server.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")))
			if err != nil {
				errlog.Println(err)
			}
		} else {
			infolog.Println("Server starting on port 8080")
			err := server.ListenAndServe(":8080")
			if err != nil {
				errlog.Println(err)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	Signals := make(chan os.Signal, 1)
	signal.Notify(Signals, os.Interrupt, syscall.SIGTERM)
	<-Signals
	infolog.Println("Server is shutting down...")
	shutdown := make(chan struct{})
	go func() {
		err = server.Shutdown()
		if err != nil {
			errlog.Println("Error during shutdown:", err)
		}
		shutdown <- struct{}{}
	}()

	select {
	case <-shutdown:
		infolog.Println("Server successfully down")
	case <-Signals:
		errlog.Println("Non-graceful shutdown")
	}
}
