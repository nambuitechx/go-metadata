package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dataModels "github.com/nambuitechx/go-metadata/models/data"
	dataServices "github.com/nambuitechx/go-metadata/services/data"
)

type DatabaseEntityHandler struct {
	DatabaseEntityService *dataServices.DatabaseEntityService
}

func InitDatabaseEntityHandler(e *gin.Engine, databaseEntityService *dataServices.DatabaseEntityService) {
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
		g.PUT("", h.createOrUpdateDatabaseEntity)
		g.DELETE("/:id", h.deleteDatabaseEntityById)
		g.DELETE("/name/:fqn", h.deleteDatabaseEntityByFqn)
	}
}

func (h *DatabaseEntityHandler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": h.DatabaseEntityService.Health() })
}

func (h *DatabaseEntityHandler) getAllDatabaseEntities(ctx *gin.Context) {
	// Get query and validate
	query := &dataModels.GetDatabaseEntitiesQuery{}

	if err := ctx.ShouldBindQuery(query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid query", "error": err.Error() })
		return
	}

	if query.Limit == 0 {
		query.Limit = 10
	}

	// Get database entites
	databaseEntities, err := h.DatabaseEntityService.GetAllDatabaseEntities(query.Service, query.Limit, query.Offset)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all databases failed", "error": err.Error() })
		return
	}

	jsonValues := []*dataModels.Database{}
	
	for _, e := range databaseEntities {
		jsonValues = append(jsonValues, e.Json)
	}

	// Get paging
	total, err := h.DatabaseEntityService.GetCountDatabaseEntities(query.Service)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all dbservices failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get all databases successfully", "data": jsonValues, "paging": total })
}

func (h *DatabaseEntityHandler) getDatabaseEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &dataModels.GetDatabaseEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	databaseEntity, err := h.DatabaseEntityService.GetDatabaseEntityById(param.ID)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Database not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, databaseEntity.Json)
}

func (h *DatabaseEntityHandler) getDatabaseEntityByFqn(ctx *gin.Context) {
	// Get param and validate
	param := &dataModels.GetDatabaseEntityByFqnParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	databaseEntity, err := h.DatabaseEntityService.GetDatabaseEntityByFqn(param.FQN)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Database not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, databaseEntity.Json)
}

func (h *DatabaseEntityHandler) createDatabaseEntity(ctx *gin.Context) {
	// Get payload
	payload := &dataModels.CreateDatabaseEntityPayload{}

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

	ctx.JSON(http.StatusCreated, databaseEntity.Json)
}

func (h *DatabaseEntityHandler) createOrUpdateDatabaseEntity(ctx *gin.Context) {
	// Get payload
	payload := &dataModels.CreateDatabaseEntityPayload{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Create or update database entity
	databaseEntity, err := h.DatabaseEntityService.CreateOrUpdateDatabaseEntity(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create or update database failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK,  databaseEntity.Json)
}

func (h *DatabaseEntityHandler) deleteDatabaseEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &dataModels.GetDatabaseEntityByIdParam{}

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
	param := &dataModels.GetDatabaseEntityByFqnParam{}

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
