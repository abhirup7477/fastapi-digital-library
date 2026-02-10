package main

import (
	"log"

	"github.com/abhirup7477/go-books-api/internal/delivery/http/handler"
	"github.com/abhirup7477/go-books-api/internal/delivery/http/router"
	"github.com/abhirup7477/go-books-api/internal/infrastructure/mailers"
	"github.com/abhirup7477/go-books-api/internal/midleware"
	"github.com/abhirup7477/go-books-api/internal/repository"
	"github.com/abhirup7477/go-books-api/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()

	repo := repository.NewInMemoryRepo()
	m, err := mailers.NewSMTPMailer()
	if err != nil {
		log.Fatalf("Environment variables loading failed: %+v\n", err)
	}
	uc := usecase.NewBookUsecase(repo, m)

	h := handler.NewHandler(uc)

	engine.Use(
		cors.New(midleware.CorsConfig()),
		midleware.ResponseTimeMidleware(),
		midleware.LoggerMidleware(),
		gin.Logger(),
		gin.Recovery(),
	)

	router.RegisterRoutes(engine, *h)

	engine.Run(":8080")
}
