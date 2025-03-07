package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nambuitechx/go-metadata/models/entities"
	"github.com/nambuitechx/go-metadata/services/entities"
)

type DatabaseSchemaEntityHandler struct {
	DatabaseSchemaEntityService *services.DatabaseSchemaEntityService
}

func InitDatabaseSchemaEntityHandler(e *gin.Engine, databaseSchemaEntityService *services.DatabaseSchemaEntityService) {
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
		g.DELETE("/:id", h.deleteDatabaseSchemaEntityById)
		g.DELETE("/name/:fqn", h.deleteDatabaseSchemaEntityByFqn)
	}
}

func (h *DatabaseSchemaEntityHandler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": h.DatabaseSchemaEntityService.Health() })
}

func (h *DatabaseSchemaEntityHandler) getAllDatabaseSchemaEntities(ctx *gin.Context) {
	// Get query and validate
	query := &models.GetDatabaseSchemaEntitiesQuery{}

	if err := ctx.ShouldBindQuery(query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid query", "error": err.Error() })
		return
	}

	if query.Limit == 0 {
		query.Limit = 10
	}

	// Get database schema entites
	databaseSchemaEntities, err := h.DatabaseSchemaEntityService.GetAllDatabaseSchemaEntities(query.Limit, query.Offset)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all database schemas failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get all database schemas successfully", "data": databaseSchemaEntities })
}

func (h *DatabaseSchemaEntityHandler) getDatabaseSchemaEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetDatabaseSchemaEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	databaseSchemaEntity, err := h.DatabaseSchemaEntityService.GetDatabaseSchemaEntityById(param.ID)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Database schema not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get database schema by id successfully", "data": databaseSchemaEntity })
}

func (h *DatabaseSchemaEntityHandler) getDatabaseSchemaEntityByFqn(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetDatabaseSchemaEntityByFqnParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	databaseSchemaEntity, err := h.DatabaseSchemaEntityService.GetDatabaseSchemaEntityByFqn(param.FQN)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Database schema not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get database schema by fqn successfully", "data": databaseSchemaEntity })
}

func (h *DatabaseSchemaEntityHandler) createDatabaseSchemaEntity(ctx *gin.Context) {
	// Get payload
	payload := &models.CreateDatabaseSchemaEntityPayload{}

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

	ctx.JSON(http.StatusCreated, gin.H{ "message": "Create database schema successfully", "data": databaseSchemaEntity })
}

func (h *DatabaseSchemaEntityHandler) deleteDatabaseSchemaEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetDatabaseSchemaEntityByIdParam{}

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
	param := &models.GetDatabaseSchemaEntityByFqnParam{}

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
