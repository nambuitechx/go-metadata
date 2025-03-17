package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	automationsModels "github.com/nambuitechx/go-metadata/models/automations"
	automationsServices "github.com/nambuitechx/go-metadata/services/automations"
)

type WorkflowEntityHandler struct {
	WorkflowEntityService *automationsServices.WorkflowEntityService
}

func InitWorkflowEntityHandler(e *gin.Engine, workflowEntityService *automationsServices.WorkflowEntityService) {
	// Init handler
	h := &WorkflowEntityHandler{ WorkflowEntityService: workflowEntityService }

	// Add routes to engine
	g := e.Group("api/v1/automations/workflows")
	{
		g.GET("/health", h.health)
		g.GET("/:id", h.getWorkflowEntityById)
		g.GET("/name/:fqn", h.getWorkflowEntityByFqn)
		g.GET("", h.getAllWorkflowEntities)
		g.POST("", h.createWorkflowEntity)
		g.POST("/trigger/:id", h.triggerWorkflowById)
		g.DELETE("/:id", h.deleteWorkflowEntityById)
		g.DELETE("/name/:fqn", h.deleteWorkflowEntityByFqn)
	}
}

func (h *WorkflowEntityHandler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": h.WorkflowEntityService.Health() })
}

func (h *WorkflowEntityHandler) getAllWorkflowEntities(ctx *gin.Context) {
	// Get query and validate
	query := &automationsModels.GetWorkflowEntitiesQuery{}

	if err := ctx.ShouldBindQuery(query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid query", "error": err.Error() })
		return
	}

	if query.Limit == 0 {
		query.Limit = 10
	}

	// Get workflow entites
	workflowEntities, err := h.WorkflowEntityService.GetAllWorkflowEntities(query.Limit, query.Offset)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all workflow failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get all workflow successfully", "data": workflowEntities })
}

func (h *WorkflowEntityHandler) getWorkflowEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &automationsModels.GetWorkflowEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	workflowEntity, err := h.WorkflowEntityService.GetWorkflowEntityById(param.ID)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Workflow not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get workflow by id successfully", "data": workflowEntity })
}

func (h *WorkflowEntityHandler) getWorkflowEntityByFqn(ctx *gin.Context) {
	// Get param and validate
	param := &automationsModels.GetWorkflowEntityByFqnParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	workflowEntity, err := h.WorkflowEntityService.GetWorkflowEntityByFqn(param.FQN)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Workflow not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get workflow by fqn successfully", "data": workflowEntity })
}

func (h *WorkflowEntityHandler) createWorkflowEntity(ctx *gin.Context) {
	// Get payload
	payload := &automationsModels.CreateWorkflowRequest{}

	if err := ctx.ShouldBindJSON(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid payload", "error": err.Error() })
		return
	}

	// Create workflow entity
	workflowEntity, err := h.WorkflowEntityService.CreateWorkflowEntity(payload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Create workflow failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{ "message": "Create workflow successfully", "data": workflowEntity })
}

func (h *WorkflowEntityHandler) triggerWorkflowById(ctx *gin.Context) {
	// Get param and validate
	param := &automationsModels.GetWorkflowEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	workflowEntity, err := h.WorkflowEntityService.GetWorkflowEntityById(param.ID)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "Workflow not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{ "message": "Trigger workflow by id successfully", "data": workflowEntity })
}

func (h *WorkflowEntityHandler) deleteWorkflowEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &automationsModels.GetWorkflowEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	err := h.WorkflowEntityService.DeleteDatabaseEntityById(param.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete workflow by id failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete workflow by id successfully" })
}

func (h *WorkflowEntityHandler) deleteWorkflowEntityByFqn(ctx *gin.Context) {
	// Get param and validate
	param := &automationsModels.GetWorkflowEntityByFqnParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	err := h.WorkflowEntityService.DeleteDatabaseEntityByFqn(param.FQN)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": "Delete workflow by fqn failed", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Delete workflow by fqn successfully" })
}
