package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dataModels "github.com/nambuitechx/go-metadata/models/data"
	dataServices "github.com/nambuitechx/go-metadata/services/data"
)

type StoredProcedureEntityHandler struct {
	StoredProcedureEntityService *dataServices.StoredProcedureEntityService
}

func InitStoreProcedureEntityHandler(e *gin.Engine, storedProcedureEntityService *dataServices.StoredProcedureEntityService) {
	// Init handler
	h := &StoredProcedureEntityHandler{ StoredProcedureEntityService: storedProcedureEntityService }

	// Add routes to engine
	g := e.Group("api/v1/storedProcedures")
	{
		g.GET("/health", h.health)
		g.GET("/:id", h.getStoredProcedureEntityById)
		g.GET("/name/:fqn", h.getStoredProcedureEntityByFqn)
		g.GET("", h.getAllStoredProcedureEntities)
		g.POST("", h.createStoredProcedureEntity)
		g.PUT("", h.createOrUpdateStoredProcedureEntity)
		g.DELETE("/:id", h.deleteStoredProcedureEntityById)
		g.DELETE("/name/:fqn", h.deleteStoredProceduredEntityByFqn)
	}
}

func (h *StoredProcedureEntityHandler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": h.StoredProcedureEntityService.Health() })
}

func (h *StoredProcedureEntityHandler) getAllStoredProcedureEntities(ctx *gin.Context) {
	// Get query and validate
	query := &dataModels.GetTableEntitiesQuery{}

	if err := ctx.ShouldBindQuery(query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid query", "error": err.Error() })
		return
	}

	if query.Limit == 0 {
		query.Limit = 10
	}

	// Get stored procedure entites
	tableEntities, err := h.StoredProcedureEntityService.GetAllStoredProcedureEntities(query.Limit, query.Offset)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all stored procedure failed", "error": err.Error() })
		return
	}

	jsonValues := []*dataModels.StoredProcedure{}
	
	for _, e := range tableEntities {
		jsonValues = append(jsonValues, e.Json)
	}

	// Get paging
	total, err := h.StoredProcedureEntityService.GetCountStoredProcedureEntities()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all stored procedure failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get all stored procedure successfully", "data": jsonValues, "paging": total })
}

func (h *StoredProcedureEntityHandler) getStoredProcedureEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &dataModels.GetStoredProcedureEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	storedProcedureEntity, err := h.StoredProcedureEntityService.GetStoredProcedureEntityById(param.ID)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Stored procedure not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, storedProcedureEntity.Json)
}

func (h *StoredProcedureEntityHandler) getStoredProcedureEntityByFqn(ctx *gin.Context) {
	// Get param and validate
	param := &dataModels.GetStoredProcedureEntityByFqnParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	storedProcedureEntity, err := h.StoredProcedureEntityService.GetStoredProcedureEntityByFqn(param.FQN)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Stored procedure not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, storedProcedureEntity.Json)
}

func (h *StoredProcedureEntityHandler) createStoredProcedureEntity(ctx *gin.Context) {
	// Get payload
	payload := &dataModels.CreateStoredProcedureEntityPayload{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Create stored procedure entity
	storedProcedureEntity, err := h.StoredProcedureEntityService.CreateStoredProcedureEntity(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create stored procedure failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusCreated, storedProcedureEntity.Json)
}

func (h *StoredProcedureEntityHandler) createOrUpdateStoredProcedureEntity(ctx *gin.Context) {
	// Get payload
	payload := &dataModels.CreateStoredProcedureEntityPayload{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Create or update stored procedure entity
	storedProcedureEntity, err := h.StoredProcedureEntityService.CreateOrUpdateStoredProcedureEntity(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create or update stored procedure failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, storedProcedureEntity.Json)
}

func (h *StoredProcedureEntityHandler) deleteStoredProcedureEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &dataModels.GetStoredProcedureEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	err := h.StoredProcedureEntityService.DeleteStoredProcedureEntityById(param.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete stored procedure by id failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete stored procedure by id successfully" })
}

func (h *StoredProcedureEntityHandler) deleteStoredProceduredEntityByFqn(ctx *gin.Context) {
	// Get param and validate
	param := &dataModels.GetStoredProcedureEntityByFqnParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	err := h.StoredProcedureEntityService.DeleteStoredProcedureEntityByFqn(param.FQN)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete stored procedure by fqn failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete stored procedure by fqn successfully" })
}
