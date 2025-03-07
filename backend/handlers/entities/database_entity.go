package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nambuitechx/go-metadata/models/entities"
	"github.com/nambuitechx/go-metadata/services/entities"
)

type DatabaseEntityHandler struct {
	DatabaseEntityService *services.DatabaseEntityService
}

func InitDatabaseEntityHandler(e *gin.Engine, databaseEntityService *services.DatabaseEntityService) {
	// Init handler
	h := &DatabaseEntityHandler{ DatabaseEntityService: databaseEntityService }

	// Add routes to engine
	g := e.Group("api/v1/databases")
	{
		g.GET("/health", h.health)
		g.GET("/:id", h.getDatabaseEntityById)
		g.GET("/name/:fqn", h.getDatabaseEntityByFqn)
		g.GET("", h.getAllDatabaseEntities)
		g.POST("", h.createDatabaseEntity)
		g.DELETE("/:id", h.deleteDatabaseEntityById)
		g.DELETE("/name/:fqn", h.deleteDatabaseEntityByFqn)
	}
}

func (h *DatabaseEntityHandler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": h.DatabaseEntityService.Health() })
}

func (h *DatabaseEntityHandler) getAllDatabaseEntities(ctx *gin.Context) {
	// Get query and validate
	query := &models.GetDatabaseEntitiesQuery{}

	if err := ctx.ShouldBindQuery(query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid query", "error": err.Error() })
		return
	}

	if query.Limit == 0 {
		query.Limit = 10
	}

	// Get database entites
	databaseEntities, err := h.DatabaseEntityService.GetAllDatabaseEntities(query.Limit, query.Offset)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all databases failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get all databases successfully", "data": databaseEntities })
}

func (h *DatabaseEntityHandler) getDatabaseEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetDatabaseEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	databaseEntity, err := h.DatabaseEntityService.GetDatabaseEntityById(param.ID)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Database not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get database by id successfully", "data": databaseEntity })
}

func (h *DatabaseEntityHandler) getDatabaseEntityByFqn(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetDatabaseEntityByFqnParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	databaseEntity, err := h.DatabaseEntityService.GetDatabaseEntityByFqn(param.FQN)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Database not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get database by fqn successfully", "data": databaseEntity })
}

func (h *DatabaseEntityHandler) createDatabaseEntity(ctx *gin.Context) {
	// Get payload
	payload := &models.CreateDatabaseEntityPayload{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Create database entity
	databaseEntity, err := h.DatabaseEntityService.CreateDatabaseEntity(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create database failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{ "message": "Create database successfully", "data": databaseEntity })
}

func (h *DatabaseEntityHandler) deleteDatabaseEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetDatabaseEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	err := h.DatabaseEntityService.DeleteDatabaseEntityById(param.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete database by id failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete database by id successfully" })
}

func (h *DatabaseEntityHandler) deleteDatabaseEntityByFqn(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetDatabaseEntityByFqnParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	err := h.DatabaseEntityService.DeleteDatabaseEntityByFqn(param.FQN)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete database by fqn failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete database by fqn successfully" })
}
