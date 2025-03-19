package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nambuitechx/go-metadata/configs"
	systemModels "github.com/nambuitechx/go-metadata/models/system"
)

type SystemHandler struct {
	SystemVersion *systemModels.SystemVersion
}

func InitDBServiceEntityHandler(e *gin.Engine, settings *configs.Settings) {
	// Setup system version
	systemVersion := &systemModels.SystemVersion{
		Version: settings.SystemVersion,
		Revision: settings.SystemRevision,
		Timestamp: settings.SystemTimestamp,
	}

	// Init handler
	h := &SystemHandler{ SystemVersion: systemVersion }

	// Add routes to engine
	g := e.Group("api/v1/system")
	{
		g.GET("/health", h.health)
		g.GET("/version", h.getCatalogVersion)
	}
}

func (h *SystemHandler) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "message": "System service is available" })
}

func (h *SystemHandler) getCatalogVersion(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, h.SystemVersion)
}
