package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nambuitechx/go-metadata/models/entities"
	"github.com/nambuitechx/go-metadata/services/entities"
)

type DBServiceEntityHandler struct {
	DBServiceEntityService *services.DBServiceEntityService
}

func InitDBServiceEntityHandler(e *gin.Engine, dbserviceEntityService *services.DBServiceEntityService) {
	// Init handler
	h := &DBServiceEntityHandler{ DBServiceEntityService: dbserviceEntityService }

	// Add routes to engine
	g := e.Group("api/v1/services/databaseServices")
	{
		g.GET("/health", h.health)
		g.GET("/:id", h.getDBServiceEntityById)
		g.GET("/name/:fqn", h.getDBServiceEntityByFqn)
		g.GET("", h.getAllDBServiceEntities)
		g.POST("", h.createDBServiceEntity)
		g.DELETE("/:id", h.deleteDBServiceEntityById)
		g.DELETE("/name/:fqn", h.deleteDBServiceEntityByFqn)
	}
}

func (h *DBServiceEntityHandler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": h.DBServiceEntityService.Health() })
}

func (h *DBServiceEntityHandler) getAllDBServiceEntities(ctx *gin.Context) {
	// Get query and validate
	query := &models.GetDBServiceEntitiesQuery{}

	if err := ctx.ShouldBindQuery(query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid query", "error": err.Error() })
		return
	}

	if query.Limit == 0 {
		query.Limit = 10
	}

	// Get dbservice entites
	dbserviceEntities, err := h.DBServiceEntityService.GetAllDBServiceEntities(query.Limit, query.Offset)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all dbservices failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get all dbservices successfully", "data": dbserviceEntities })
}

func (h *DBServiceEntityHandler) getDBServiceEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetDBServiceEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	dbserviceEntity, err := h.DBServiceEntityService.GetDBServiceEntityById(param.ID)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "DBService not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get dbservice by id successfully", "data": dbserviceEntity })
}

func (h *DBServiceEntityHandler) getDBServiceEntityByFqn(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetDBServiceEntityByFqnParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	dbserviceEntity, err := h.DBServiceEntityService.GetDBServiceEntityByFqn(param.FQN)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "DBService not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get dbservice by fqn successfully", "data": dbserviceEntity })
}

func (h *DBServiceEntityHandler) createDBServiceEntity(ctx *gin.Context) {
	// Get payload
	payload := &models.CreateDBServiceEntityPayload{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Validate payload
	if err := models.ValidateCreateDBServiceEntityPayload(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create dbservice failed", "error": err.Error() })
		return
	}

	// Create dbservice entity
	dbserviceEntity, err := h.DBServiceEntityService.CreateDBServiceEntity(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create dbservice failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{ "message": "Create dbservice successfully", "data": dbserviceEntity })
}

func (h *DBServiceEntityHandler) deleteDBServiceEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetDBServiceEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	err := h.DBServiceEntityService.DeleteDBServiceEntityById(param.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete dbservice by id failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete dbservice by id successfully" })
}

func (h *DBServiceEntityHandler) deleteDBServiceEntityByFqn(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetDBServiceEntityByFqnParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	err := h.DBServiceEntityService.DeleteDBServiceEntityByFqn(param.FQN)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete dbservice by fqn failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete dbservice by fqn successfully" })
}
