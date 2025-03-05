package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nambuitechx/go-metadata/models"
	"github.com/nambuitechx/go-metadata/services"
)

type DBServiceHandler struct {
	DBServiceService *services.DBServiceService
}

func InitDBServiceHandler(e *gin.Engine, dbserviceService *services.DBServiceService) {
	// Init handler
	h := &DBServiceHandler{ DBServiceService: dbserviceService }

	// Add routes to engine
	g := e.Group("api/v1/dbservices")
	{
		g.GET("/health", h.health)
		g.GET("/:id", h.getDBServiceEntityById)
		g.GET("", h.getAllDBServiceEntities)
		g.POST("", h.createDBServiceEntity)
		g.DELETE("/:id", h.deleteDBServiceEntityById)
	}
}

func (h *DBServiceHandler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": h.DBServiceService.Health() })
}

func (h *DBServiceHandler) getAllDBServiceEntities(ctx *gin.Context) {
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
	dbserviceEntities, err := h.DBServiceService.GetAllDBServiceEntities(query.Limit, query.Offset)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all dbservices failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get all dbservices successfully", "data": dbserviceEntities })
}

func (h *DBServiceHandler) getDBServiceEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetDBServiceEntityParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	// Get dbservice entity
	dbserviceEntity, err := h.DBServiceService.GetDBServiceEntityById(param.ID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "DBService entity not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get dbservice entity by id successfully", "data": dbserviceEntity })
}

func (h *DBServiceHandler) createDBServiceEntity(ctx *gin.Context) {
	// Get payload
	payload := &models.CreateDBServiceEntity{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Validate payload
	if err := models.ValidateCreateDBServiceEntityPayload(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create dbservice entity failed", "error": err.Error() })
		return
	}

	// Create dbservice entity
	dbserviceEntity, err := h.DBServiceService.CreateDBServiceEntity(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create dbservice entity failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Create dbservice entity successfully", "data": dbserviceEntity })
}

func (h *DBServiceHandler) deleteDBServiceEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &models.GetDBServiceEntityParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	// Delete dbservice entity
	err := h.DBServiceService.DeleteDBServiceEntityById(param.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete dbservice entity by id failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete dbservice entity by id successfully" })
}
