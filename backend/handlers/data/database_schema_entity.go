package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dataModels "github.com/nambuitechx/go-metadata/models/data"
	dataServices "github.com/nambuitechx/go-metadata/services/data"
)

type DatabaseSchemaEntityHandler struct {
	DatabaseSchemaEntityService *dataServices.DatabaseSchemaEntityService
}

func InitDatabaseSchemaEntityHandler(e *gin.Engine, databaseSchemaEntityService *dataServices.DatabaseSchemaEntityService) {
	// Init handler
	h := &DatabaseSchemaEntityHandler{ DatabaseSchemaEntityService: databaseSchemaEntityService }

	// Add routes to engine
	g := e.Group("api/v1/databaseSchemas")
	{
		g.GET("/health", h.health)
		g.GET("/:id", h.getDatabaseSchemaEntityById)
		g.GET("/name/:fqn", h.getDatabaseSchemaEntityByFqn)
		g.GET("", h.getAllDatabaseSchemaEntities)
		g.POST("", h.createDatabaseSchemaEntity)
		g.PUT("", h.createOrUpdateDatabaseSchemaEntity)
		g.DELETE("/:id", h.deleteDatabaseSchemaEntityById)
		g.DELETE("/name/:fqn", h.deleteDatabaseSchemaEntityByFqn)
	}
}

func (h *DatabaseSchemaEntityHandler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": h.DatabaseSchemaEntityService.Health() })
}

func (h *DatabaseSchemaEntityHandler) getAllDatabaseSchemaEntities(ctx *gin.Context) {
	// Get query and validate
	query := &dataModels.GetDatabaseSchemaEntitiesQuery{}

	if err := ctx.ShouldBindQuery(query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid query", "error": err.Error() })
		return
	}

	if query.Limit == 0 {
		query.Limit = 10
	}

	// Get database schema entites
	databaseSchemaEntities, err := h.DatabaseSchemaEntityService.GetAllDatabaseSchemaEntities(query.Database, query.Limit, query.Offset)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all database schemas failed", "error": err.Error() })
		return
	}

	jsonValues := []*dataModels.DatabaseSchema{}
	
	for _, e := range databaseSchemaEntities {
		jsonValues = append(jsonValues, e.Json)
	}

	// Get paging
	total, err := h.DatabaseSchemaEntityService.GetCountDatabaseSchemaEntities(query.Database)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all dbservices failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get all database schemas successfully", "data": jsonValues, "paging": total })
}

func (h *DatabaseSchemaEntityHandler) getDatabaseSchemaEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &dataModels.GetDatabaseSchemaEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	databaseSchemaEntity, err := h.DatabaseSchemaEntityService.GetDatabaseSchemaEntityById(param.ID)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Database schema not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, databaseSchemaEntity.Json)
}

func (h *DatabaseSchemaEntityHandler) getDatabaseSchemaEntityByFqn(ctx *gin.Context) {
	// Get param and validate
	param := &dataModels.GetDatabaseSchemaEntityByFqnParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	databaseSchemaEntity, err := h.DatabaseSchemaEntityService.GetDatabaseSchemaEntityByFqn(param.FQN)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Database schema not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, databaseSchemaEntity.Json)
}

func (h *DatabaseSchemaEntityHandler) createDatabaseSchemaEntity(ctx *gin.Context) {
	// Get payload
	payload := &dataModels.CreateDatabaseSchemaEntityPayload{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Create database schema entity
	databaseSchemaEntity, err := h.DatabaseSchemaEntityService.CreateDatabaseSchemaEntity(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create database schema failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusCreated, databaseSchemaEntity.Json)
}

func (h *DatabaseSchemaEntityHandler) createOrUpdateDatabaseSchemaEntity(ctx *gin.Context) {
	// Get payload
	payload := &dataModels.CreateDatabaseSchemaEntityPayload{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Create or update database schema entity
	databaseSchemaEntity, err := h.DatabaseSchemaEntityService.CreateOrUpdateDatabaseSchemaEntity(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create or update database schema failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, databaseSchemaEntity.Json)
}

func (h *DatabaseSchemaEntityHandler) deleteDatabaseSchemaEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &dataModels.GetDatabaseSchemaEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	err := h.DatabaseSchemaEntityService.DeleteDatabaseSchemaEntityById(param.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete database schema by id failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete database schema by id successfully" })
}

func (h *DatabaseSchemaEntityHandler) deleteDatabaseSchemaEntityByFqn(ctx *gin.Context) {
	// Get param and validate
	param := &dataModels.GetDatabaseSchemaEntityByFqnParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	err := h.DatabaseSchemaEntityService.DeleteDatabaseSchemaEntityByFqn(param.FQN)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete database schema by fqn failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete database schema by fqn successfully" })
}
