package main

import (
	"APPOINMENT_BOOKING_SYSTEM/config"
	mysql "APPOINMENT_BOOKING_SYSTEM/db"
	"APPOINMENT_BOOKING_SYSTEM/routes"
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	config.GetConfigurations()
	mysql.ConnectMySQL()
}

func main() {

	defer func() {
		log.Println("Closing Mysql database connection")
		mysql.CloseMySQL()
	}()

	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	router := gin.New()
	routes.RegisterRouter(router)

	server := &http.Server{
		Addr:         ":" + config.AppConfig.Server.Port,
		BaseContext:  func(listener net.Listener) context.Context { return ctx },
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}
	errorChan := make(chan error, 1)

	go func() {
		errorChan <- server.ListenAndServe()
	}()
	log.Printf("server is up and running on hhtp://localhost:::%s", config.AppConfig.Server.Port)

	select {
	case err := <-errorChan:
		return err
	case <-ctx.Done():
		log.Println("Server interrupt received, shutting down ...")
		cancel()
		return server.Shutdown(context.Background())
	}
}
