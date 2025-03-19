package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	servicesModels "github.com/nambuitechx/go-metadata/models/services"
	servicesServices "github.com/nambuitechx/go-metadata/services/services"
)

type TestConnectionDefinitionEntityHandler struct {
	TestConnectionDefinitionEntityService *servicesServices.TestConnectionDefinitionEntityService
}

func InitTestConnectionDefinitionEntityHandler(e *gin.Engine, testConnectionDefinitionEntityService *servicesServices.TestConnectionDefinitionEntityService) {
	// Init handler
	h := &TestConnectionDefinitionEntityHandler{ TestConnectionDefinitionEntityService: testConnectionDefinitionEntityService }

	// Add routes to engine
	g := e.Group("api/v1/services/testConnectionDefinitions")
	{
		g.GET("/health", h.health)
		g.GET("/:id", h.getTestConnectionDefinitionEntityById)
		g.GET("/name/:fqn", h.getTestConnectionDefinitionEntityByFqn)
		g.GET("", h.getAllTestConnectionDefinitionEntities)
	}
}

func (h *TestConnectionDefinitionEntityHandler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": h.TestConnectionDefinitionEntityService.Health() })
}

func (h *TestConnectionDefinitionEntityHandler) getAllTestConnectionDefinitionEntities(ctx *gin.Context) {
	// Get query and validate
	query := &servicesModels.GetTestConnectionDefinitionEntitiesQuery{}

	if err := ctx.ShouldBindQuery(query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid query", "error": err.Error() })
		return
	}

	if query.Limit == 0 {
		query.Limit = 10
	}

	// Get test connection definition entites
	testConnectionDefinitionEntities, err := h.TestConnectionDefinitionEntityService.GetAllTestConnectionDefinitionEntities(query.Limit, query.Offset)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Get all test connection definition failed", "error": err.Error() })
		return
	}

	jsonValues := []*servicesModels.TestConnectionDefinition{}
	
	for _, e := range testConnectionDefinitionEntities {
		jsonValues = append(jsonValues, e.Json)
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Get all test connection definition successfully", "data": jsonValues })
}

func (h *TestConnectionDefinitionEntityHandler) getTestConnectionDefinitionEntityById(ctx *gin.Context) {
	// Get param and validate
	param := &servicesModels.GetTestConnectionDefinitionEntityByIdParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	testConnectionDefinitionEntity, err := h.TestConnectionDefinitionEntityService.GetTestConnectionDefinitionEntityById(param.ID)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "TestConnectionDefinition not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, testConnectionDefinitionEntity.Json)
}

func (h *TestConnectionDefinitionEntityHandler) getTestConnectionDefinitionEntityByFqn(ctx *gin.Context) {
	// Get param and validate
	param := &servicesModels.GetTestConnectionDefinitionEntityByFqnParam{}

	if err := ctx.ShouldBindUri(param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid param", "error": err.Error() })
		return
	}

	testConnectionDefinitionEntity, err := h.TestConnectionDefinitionEntityService.GetTestConnectionDefinitionEntityByFqn(param.FQN)
	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{ "message": "TestConnectionDefinition not found", "error": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, testConnectionDefinitionEntity.Json )
}
