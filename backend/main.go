package main

import (
	"flag"
	"flexio-api/config"
	"flexio-api/internal/app"
	"flexio-api/internal/routes"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "port to listen on")
	flag.Parse()

	// Load the config
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	app, err := app.NewApplication(cfg)
	if err != nil {
		panic(err)
	}

	defer app.DB.Close()

	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
	}

	app.Logger.Println("Flexio API is running")

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
