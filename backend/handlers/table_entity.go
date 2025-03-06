package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nambuitechx/go-metadata/models"
	"github.com/nambuitechx/go-metadata/services"
)

type TableEntityHandler struct {
	TableEntityService *services.TableEntityService
}

func InitTableEntityHandler(e *gin.Engine, tableEntityService *services.TableEntityService) {
	// Init handler
	h := &TableEntityHandler{ TableEntityService: tableEntityService }

	// Add routes to engine
	g := e.Group("api/v1/tables")
	{
		g.GET("/health", h.health)
		g.GET("/:id", h.getTableEntityById)
		g.GET("", h.getAllTableEntities)
		g.POST("", h.createTableEntity)
		g.DELETE("/:id", h.deleteTableEntityById)
	}
}

func (h *TableEntityHandler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": h.TableEntityService.Health() })
}

func (h *TableEntityHandler) getAllTableEntities(ctx *gin.Context) {
	// Get query and validate
	query := &models.GetTableEntitiesQuery{}

	if err := ctx.ShouldBindQuery(query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid query", "error": err.Error() })
		return
	}

	if query.Limit == 0 {
		query.Limit = 10
	}

	// Get table entites
	tableEntities, err := h.TableEntityService.GetAllTableEntities(query.Limit, query.Offset)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all table failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get all table successfully", "data": tableEntities })
}

func (h *TableEntityHandler) getTableEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetTableEntityParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	var tableEntity *models.TableEntity
	var err error

	// Get table entity by either uuid or fully qualified name
	if invalid := uuid.Validate(param.ID); invalid == nil {
		tableEntity, err = h.TableEntityService.GetTableEntityById(param.ID)
	} else {
		tableEntity, err = h.TableEntityService.GetTableEntityByFqn(param.ID)
	}
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Table not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get table by id or fqn successfully", "data": tableEntity })
}

func (h *TableEntityHandler) createTableEntity(ctx *gin.Context) {
	// Get payload
	payload := &models.CreateTableEntityPayload{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Create table entity
	tableEntity, err := h.TableEntityService.CreateTableEntity(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create table failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Create table successfully", "data": tableEntity })
}

func (h *TableEntityHandler) deleteTableEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetTableEntityParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	var err error

	// Delete table by either uuid or fully qualified name
	if invalid := uuid.Validate(param.ID); invalid == nil {
		err = h.TableEntityService.DeleteTableEntityById(param.ID)
	} else {
		err = h.TableEntityService.DeleteTableEntityByFqn(param.ID)
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete table by id or fqn failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete table by id or fqn successfully" })
}
