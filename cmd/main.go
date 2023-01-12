package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/juajosserand/goweb-challenge/cmd/handler"
	"github.com/juajosserand/goweb-challenge/internal/domain"
	"github.com/juajosserand/goweb-challenge/internal/ticket"
	"github.com/juajosserand/goweb-challenge/pkg/httpserver"
	"github.com/juajosserand/goweb-challenge/pkg/storage"
)

func main() {
	// load env
	if err := godotenv.Load(); err != nil {
		log.Println(err)
		panic("could not load env")
	}

	// load tickets from file
	var tickets []domain.Ticket
	if err := storage.ReadCSV(os.Getenv("FILEPATH"), &tickets); err != nil {
		log.Println(err)
		panic("could not load tickets")
	}

	// repository
	repo := ticket.NewRepository(tickets)

	// service
	svc := ticket.NewService(repo)

	// htte server
	mux := gin.Default()
	handler.NewService(mux, svc)
	server := httpserver.New(mux, httpserver.Port(os.Getenv("PORT")))

	// signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("signal:", s.String())
	case err := <-server.Notify():
		log.Println(fmt.Errorf("error: %w", err))
	}

	if err := server.Shutdown(); err != nil {
		log.Println(fmt.Errorf("error: %w", err))
	}
}
