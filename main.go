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

	// Create a new server
	server := &fasthttp.Server{
		ReadTimeout:     time.Second * 30,
		WriteTimeout:    time.Second * 30,
		CloseOnShutdown: true,
		Handler:         func(ctx *fasthttp.RequestCtx) {},
	}

	r := routes()
	server.Handler = r.Handler

	go func() {
		if os.Getenv("PORT") != "" {
			infolog.Println("Server starting on port", os.Getenv("PORT"))
			errlog.Fatalln(server.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT"))))
		} else {
			infolog.Println("Server starting on port 8080")
			errlog.Fatalln(server.ListenAndServe(":8080"))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	infolog.Println("Server is shutting down...")
	err = server.Shutdown()
	if err != nil {
		errlog.Println("Error during shutdown:", err)
	}
}
