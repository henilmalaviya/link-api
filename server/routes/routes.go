package routes

import (
	v1Handlers "api.link.henil.dev/server/handlers/v1"
	"api.link.henil.dev/server/middlewares"
	"github.com/gin-gonic/gin"
)

func registerV1WorkspaceRoutes(rg *gin.RouterGroup) {
	workspace := rg.Group("/workspaces")
	{
		workspace.POST("/", v1Handlers.CreateWorkspace)

		protected := workspace.Group("/")
		protected.Use(middlewares.WorkspaceSecretAuthMiddleware())
		{
			protected.GET("/:id", v1Handlers.GetWorkspaceByID)
			protected.GET("/:id/stats", v1Handlers.GetWorkspaceStatsByID)
		}
	}
}

func registerV1ApiKeyRoutes(rg *gin.RouterGroup) {
	apiKey := rg.Group("/api-keys")
	apiKey.Use(middlewares.WorkspaceAuthMiddleware())
	{
		apiKey.POST("/", v1Handlers.CreateApiKey)
		apiKey.GET("/", v1Handlers.GetAllApiKeys)
	}
}

func registerV1LinksRoutes(rg *gin.RouterGroup) {
	// links := rg.Group("/links")
	// {
	// }
}

func Register(r *gin.Engine) {

	api := r.Group("/api")

	{
		v1 := api.Group("/v1")
		{
			registerV1WorkspaceRoutes(v1)
			registerV1ApiKeyRoutes(v1)
			registerV1LinksRoutes(v1)
		}
	}

}
