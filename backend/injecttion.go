package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nambuitechx/go-metadata/configs"
	"github.com/nambuitechx/go-metadata/handlers"
	"github.com/nambuitechx/go-metadata/repositories"
	"github.com/nambuitechx/go-metadata/services"
)

func getEngine() *gin.Engine {
	// Connect to database
	settings := configs.NewSettings()
	db := configs.NewDatabaseConnection(settings).DB

	// Repositories
	dbserviceEntityRepository := repositories.NewDBServiceEntityRepository(db)

	// Services
	dbserviceService := services.NewDBServiceService(dbserviceEntityRepository)

	// Engine
	engine := gin.Default()

	// Middlewares
	config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
    config.ExposeHeaders = []string{"Content-Length"}
    config.AllowCredentials = true

	engine.Use(cors.New(config))

	// Routes
	engine.GET("/health", checkHealth)
	handlers.InitDBServiceHandler(engine, dbserviceService)

	return engine
}

func checkHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H { "message": "Healthy" })
}
