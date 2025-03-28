package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dataModels "github.com/nambuitechx/go-metadata/models/data"
	dataServices "github.com/nambuitechx/go-metadata/services/data"
)

type TableEntityHandler struct {
	TableEntityService *dataServices.TableEntityService
}

func InitTableEntityHandler(e *gin.Engine, tableEntityService *dataServices.TableEntityService) {
	// Init handler
	h := &TableEntityHandler{ TableEntityService: tableEntityService }

	// Add routes to engine
	g := e.Group("api/v1/tables")
	{
		g.GET("/health", h.health)
		g.GET("/:id", h.getTableEntityById)
		g.GET("/name/:fqn", h.getTableEntityByFqn)
		g.GET("", h.getAllTableEntities)
		g.POST("", h.createTableEntity)
		g.PUT("", h.createOrUpdateTableEntity)
		g.DELETE("/:id", h.deleteTableEntityById)
		g.DELETE("/name/:fqn", h.deleteTableEntityByFqn)
	}
}

func (h *TableEntityHandler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": h.TableEntityService.Health() })
}

func (h *TableEntityHandler) getAllTableEntities(ctx *gin.Context) {
	// Get query and validate
	query := &dataModels.GetTableEntitiesQuery{}

	if err := ctx.ShouldBindQuery(query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid query", "error": err.Error() })
		return
	}

	if query.Limit == 0 {
		query.Limit = 10
	}

	// Get table entites
	tableEntities, err := h.TableEntityService.GetAllTableEntities(query.DatabaseSchema, query.Limit, query.Offset)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all table failed", "error": err.Error() })
		return
	}

	jsonValues := []*dataModels.Table{}
	
	for _, e := range tableEntities {
		jsonValues = append(jsonValues, e.Json)
	}

	// Get paging
	total, err := h.TableEntityService.GetCountTableEntities(query.DatabaseSchema)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all dbservices failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get all table successfully", "data": jsonValues, "paging": total })
}

func (h *TableEntityHandler) getTableEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &dataModels.GetTableEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	tableEntity, err := h.TableEntityService.GetTableEntityById(param.ID)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Table not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, tableEntity.Json)
}

func (h *TableEntityHandler) getTableEntityByFqn(ctx *gin.Context) {
	// Get param and validate
	param := &dataModels.GetTableEntityByFqnParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	tableEntity, err := h.TableEntityService.GetTableEntityByFqn(param.FQN)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Table not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, tableEntity.Json)
}

func (h *TableEntityHandler) createTableEntity(ctx *gin.Context) {
	// Get payload
	payload := &dataModels.CreateTableEntityPayload{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Validate payload
	if err := dataModels.ValidateCreateTableEntityPayload(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create table failed", "error": err.Error() })
		return
	}

	// Create table entity
	tableEntity, err := h.TableEntityService.CreateTableEntity(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create table failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusCreated, tableEntity.Json)
}

func (h *TableEntityHandler) createOrUpdateTableEntity(ctx *gin.Context) {
	// Get payload
	payload := &dataModels.CreateTableEntityPayload{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Validate payload
	if err := dataModels.ValidateCreateTableEntityPayload(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create table failed", "error": err.Error() })
		return
	}

	// Create or update table entity
	tableEntity, err := h.TableEntityService.CreateOrUpdateTableEntity(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create or update table failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, tableEntity.Json)
}

func (h *TableEntityHandler) deleteTableEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &dataModels.GetTableEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	err := h.TableEntityService.DeleteTableEntityById(param.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete table by id failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete table by id successfully" })
}

func (h *TableEntityHandler) deleteTableEntityByFqn(ctx *gin.Context) {
	// Get param and validate
	param := &dataModels.GetTableEntityByFqnParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	err := h.TableEntityService.DeleteTableEntityByFqn(param.FQN)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete table by fqn failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete table by fqn successfully" })
}
