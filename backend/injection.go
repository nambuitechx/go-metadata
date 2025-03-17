package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nambuitechx/go-metadata/configs"
	servicesHandlers "github.com/nambuitechx/go-metadata/handlers/services"
	dataHandlers "github.com/nambuitechx/go-metadata/handlers/data"
	servicesServices "github.com/nambuitechx/go-metadata/services/services"
	dataServices "github.com/nambuitechx/go-metadata/services/data"
	servicesRepositories "github.com/nambuitechx/go-metadata/repositories/services"
	dataRepositories "github.com/nambuitechx/go-metadata/repositories/data"
)

func getEngine() *gin.Engine {
	// Connect to database
	settings := configs.NewSettings()
	db := configs.NewDatabaseConnection(settings).DB

	// Repositories
	dbserviceEntityRepository := servicesRepositories.NewDBServiceEntityRepository(db)
	databaseEntityRepository := dataRepositories.NewDatabaseEntityRepository(db)
	databaseSchemaEntityRepository := dataRepositories.NewDatabaseSchemaEntityRepository(db)
	tableEntityRepository := dataRepositories.NewTableEntityRepository(db)

	// Services
	dbserviceEntityService := servicesServices.NewDBServiceEntityService(dbserviceEntityRepository)
	databaseEntityService := dataServices.NewDatabaseEntityService(dbserviceEntityRepository, databaseEntityRepository)
	databaseSchemaEntityService := dataServices.NewDatabaseSchemaEntityService(dbserviceEntityRepository, databaseEntityRepository, databaseSchemaEntityRepository)
	tableEntityService := dataServices.NewTableEntityService(dbserviceEntityRepository, databaseEntityRepository, databaseSchemaEntityRepository, tableEntityRepository)

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
	servicesHandlers.InitDBServiceEntityHandler(engine, dbserviceEntityService)
	dataHandlers.InitDatabaseEntityHandler(engine, databaseEntityService)
	dataHandlers.InitDatabaseSchemaEntityHandler(engine, databaseSchemaEntityService)
	dataHandlers.InitTableEntityHandler(engine, tableEntityService)

	return engine
}

func checkHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H { "message": "Healthy" })
}
