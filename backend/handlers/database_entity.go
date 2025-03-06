package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nambuitechx/go-metadata/models"
	"github.com/nambuitechx/go-metadata/services"
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
		g.GET("", h.getAllDatabaseEntities)
		g.POST("", h.createDatabaseEntity)
		g.DELETE("/:id", h.deleteDatabaseEntityById)
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
	param := &models.GetDatabaseEntityParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	var databaseEntity *models.DatabaseEntity
	var err error

	// Get database entity by either uuid or fully qualified name
	if invalid := uuid.Validate(param.ID); invalid == nil {
		databaseEntity, err = h.DatabaseEntityService.GetDatabaseEntityById(param.ID)
	} else {
		databaseEntity, err = h.DatabaseEntityService.GetDatabaseEntityByFqn(param.ID)
	}
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Database not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get database by id or fqn successfully", "data": databaseEntity })
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

	ctx.JSON(http.StatusOK, gin.H{ "message": "Create database successfully", "data": databaseEntity })
}

func (h *DatabaseEntityHandler) deleteDatabaseEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetDatabaseEntityParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	var err error

	// Delete database by either uuid or fully qualified name
	if invalid := uuid.Validate(param.ID); invalid == nil {
		err = h.DatabaseEntityService.DeleteDatabaseEntityById(param.ID)
	} else {
		err = h.DatabaseEntityService.DeleteDatabaseEntityByFqn(param.ID)
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete database by id or fqn failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete database by id or fqn successfully" })
}
