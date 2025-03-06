package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nambuitechx/go-metadata/configs"
	"github.com/nambuitechx/go-metadata/handlers/entities"
	"github.com/nambuitechx/go-metadata/repositories/entities"
	"github.com/nambuitechx/go-metadata/services/entities"
)

func getEngine() *gin.Engine {
	// Connect to database
	settings := configs.NewSettings()
	db := configs.NewDatabaseConnection(settings).DB

	// Repositories
	dbserviceEntityRepository := repositories.NewDBServiceEntityRepository(db)
	databaseEntityRepository := repositories.NewDatabaseEntityRepository(db)
	databaseSchemaEntityRepository := repositories.NewDatabaseSchemaEntityRepository(db)
	tableEntityRepository := repositories.NewTableEntityRepository(db)

	// Services
	dbserviceEntityService := services.NewDBServiceEntityService(dbserviceEntityRepository)
	databaseEntityService := services.NewDatabaseEntityService(dbserviceEntityRepository, databaseEntityRepository)
	databaseSchemaEntityService := services.NewDatabaseSchemaEntityService(dbserviceEntityRepository, databaseEntityRepository, databaseSchemaEntityRepository)
	tableEntityService := services.NewTableEntityService(dbserviceEntityRepository, databaseEntityRepository, databaseSchemaEntityRepository, tableEntityRepository)

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
	handlers.InitDBServiceEntityHandler(engine, dbserviceEntityService)
	handlers.InitDatabaseEntityHandler(engine, databaseEntityService)
	handlers.InitDatabaseSchemaEntityHandler(engine, databaseSchemaEntityService)
	handlers.InitTableEntityHandler(engine, tableEntityService)

	return engine
}

func checkHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H { "message": "Healthy" })
}
