package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	servicesModels "github.com/nambuitechx/go-metadata/models/services"
	servicesServices "github.com/nambuitechx/go-metadata/services/services"
)

type DBServiceEntityHandler struct {
	DBServiceEntityService *servicesServices.DBServiceEntityService
}

func InitDBServiceEntityHandler(e *gin.Engine, dbserviceEntityService *servicesServices.DBServiceEntityService) {
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
		g.PUT("", h.createOrUpdateDBServiceEntity)
		g.PUT("/:id/testConnectionResult", h.updateTestConnectionResult)
		g.DELETE("/:id", h.deleteDBServiceEntityById)
		g.DELETE("/name/:fqn", h.deleteDBServiceEntityByFqn)
	}
}

func (h *DBServiceEntityHandler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": h.DBServiceEntityService.Health() })
}

func (h *DBServiceEntityHandler) getAllDBServiceEntities(ctx *gin.Context) {
	// Get query and validate
	query := &servicesModels.GetDBServiceEntitiesQuery{}

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

	jsonValues := []*servicesModels.DBService{}
	
	for _, e := range dbserviceEntities {
		jsonValues = append(jsonValues, e.Json)
	}

	// Get paging
	total, err := h.DBServiceEntityService.GetCountDBServiceEntities()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all dbservices failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get all dbservices successfully", "data": jsonValues, "paging": total })
}

func (h *DBServiceEntityHandler) getDBServiceEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &servicesModels.GetDBServiceEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	dbserviceEntity, err := h.DBServiceEntityService.GetDBServiceEntityById(param.ID)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "DBService not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, dbserviceEntity.Json)
}

func (h *DBServiceEntityHandler) getDBServiceEntityByFqn(ctx *gin.Context) {
	// Get param and validate
	param := &servicesModels.GetDBServiceEntityByFqnParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	dbserviceEntity, err := h.DBServiceEntityService.GetDBServiceEntityByFqn(param.FQN)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "DBService not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, dbserviceEntity.Json)
}

func (h *DBServiceEntityHandler) createDBServiceEntity(ctx *gin.Context) {
	// Get payload
	payload := &servicesModels.CreateDBServiceEntityPayload{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Validate payload
	if err := servicesModels.ValidateCreateDBServiceEntityPayload(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create dbservice failed", "error": err.Error() })
		return
	}

	// Create dbservice entity
	dbserviceEntity, err := h.DBServiceEntityService.CreateDBServiceEntity(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create dbservice failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusCreated, dbserviceEntity.Json)
}

func (h *DBServiceEntityHandler) createOrUpdateDBServiceEntity(ctx *gin.Context) {
	// Get payload
	payload := &servicesModels.CreateDBServiceEntityPayload{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Validate payload
	if err := servicesModels.ValidateCreateDBServiceEntityPayload(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create dbservice failed", "error": err.Error() })
		return
	}

	// Create or update dbservice entity
	dbserviceEntity, err := h.DBServiceEntityService.CreateOrUpdateDBServiceEntity(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create or update dbservice failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, dbserviceEntity.Json)
}

func (h *DBServiceEntityHandler) updateTestConnectionResult(ctx *gin.Context) {
	// Get param and validate
	param := &servicesModels.GetDBServiceEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	// Get payload
	payload := &servicesModels.TestConnectionResult{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	dbserviceEntity, err := h.DBServiceEntityService.GetDBServiceEntityById(param.ID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "DBService not found", "error": err.Error() })
		return
	}

	dbserviceEntity.Json.TestConnectionResult = payload
	updatedDBServiceEntity, err := h.DBServiceEntityService.DBServiceEntityRepository.UpdateDBServiceEntity(dbserviceEntity)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Failed to update test connection status for dbservice", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, updatedDBServiceEntity.Json)
}

func (h *DBServiceEntityHandler) deleteDBServiceEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &servicesModels.GetDBServiceEntityByIdParam{}

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
	param := &servicesModels.GetDBServiceEntityByFqnParam{}

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
