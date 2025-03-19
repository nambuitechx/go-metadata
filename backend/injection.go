package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nambuitechx/go-metadata/configs"
	
	systemHandlers "github.com/nambuitechx/go-metadata/handlers/system"
	servicesHandlers "github.com/nambuitechx/go-metadata/handlers/services"
	dataHandlers "github.com/nambuitechx/go-metadata/handlers/data"
	automationsHandlers "github.com/nambuitechx/go-metadata/handlers/automations"
	servicesServices "github.com/nambuitechx/go-metadata/services/services"
	dataServices "github.com/nambuitechx/go-metadata/services/data"
	automationsServices "github.com/nambuitechx/go-metadata/services/automations"
	servicesRepositories "github.com/nambuitechx/go-metadata/repositories/services"
	dataRepositories "github.com/nambuitechx/go-metadata/repositories/data"
	automationsRepositories "github.com/nambuitechx/go-metadata/repositories/automations"
)

func getEngine() *gin.Engine {
	// Connect to database
	settings := configs.NewSettings()
	db := configs.NewDatabaseConnection(settings).DB

	// Repositories
	testConnectionDefinitionEntityRepository := servicesRepositories.NewTestConnectionDefinitionEntityRepository(db)
	dbserviceEntityRepository := servicesRepositories.NewDBServiceEntityRepository(db)
	databaseEntityRepository := dataRepositories.NewDatabaseEntityRepository(db)
	databaseSchemaEntityRepository := dataRepositories.NewDatabaseSchemaEntityRepository(db)
	tableEntityRepository := dataRepositories.NewTableEntityRepository(db)
	workflowEntityRepository := automationsRepositories.NewWorkflowEntityRepository(db)

	// Services
	testConnectionDefinitionEntityService := servicesServices.NewTestConnectionDefinitionEntityService(testConnectionDefinitionEntityRepository)
	dbserviceEntityService := servicesServices.NewDBServiceEntityService(dbserviceEntityRepository)
	databaseEntityService := dataServices.NewDatabaseEntityService(dbserviceEntityRepository, databaseEntityRepository)
	databaseSchemaEntityService := dataServices.NewDatabaseSchemaEntityService(dbserviceEntityRepository, databaseEntityRepository, databaseSchemaEntityRepository)
	tableEntityService := dataServices.NewTableEntityService(dbserviceEntityRepository, databaseEntityRepository, databaseSchemaEntityRepository, tableEntityRepository)
	workflowEntityService := automationsServices.NewWorkflowEntityService(workflowEntityRepository)

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
	systemHandlers.InitDBServiceEntityHandler(engine, settings)
	servicesHandlers.InitTestConnectionDefinitionEntityHandler(engine, testConnectionDefinitionEntityService)
	servicesHandlers.InitDBServiceEntityHandler(engine, dbserviceEntityService)
	dataHandlers.InitDatabaseEntityHandler(engine, databaseEntityService)
	dataHandlers.InitDatabaseSchemaEntityHandler(engine, databaseSchemaEntityService)
	dataHandlers.InitTableEntityHandler(engine, tableEntityService)
	automationsHandlers.InitWorkflowEntityHandler(engine, workflowEntityService)

	return engine
}

func checkHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H { "message": "Healthy" })
}
