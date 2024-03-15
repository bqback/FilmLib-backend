package main

import (
	"context"
	"filmlib/internal/config"
	"filmlib/internal/handlers"
	"filmlib/internal/logging"
	"filmlib/internal/mux"
	"filmlib/internal/service"
	"filmlib/internal/storage"
	"filmlib/internal/storage/postgresql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

const configPath string = "config/config.yml"
const envPath string = "config/.env"

func main() {
	config, err := config.LoadConfig(envPath, configPath)
	// govalidator.SetFieldsRequiredByDefault(true)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Config loaded")

	logger, err := logging.NewLogrusLogger(config.Logging)
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("Logger configured")

	dbConnection, err := postgresql.GetDBConnection(*config.Database)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer dbConnection.Close()
	logger.Info("Database connection established")

	storages := storage.NewPostgresStorages(dbConnection)
	logger.Info("Storages configured")

	services := service.NewServices(storages)
	logger.Info("Services configured")

	handlers := handlers.NewHandlers(services, config)
	logger.Info("Handlers configured")

	mux := mux.SetupMux(handlers, config, &logger)
	logger.Info("Router configured")

	var server = http.Server{
		Addr:    fmt.Sprintf(":%d", config.Server.Port),
		Handler: *mux,
	}

	logger.Info("Server is running")

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			logger.Info("HTTP server Shutdown: " + err.Error())
		}
		close(idleConnsClosed)
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		logger.Fatal("HTTP server ListenAndServe: " + err.Error())
	}

	<-idleConnsClosed
}
