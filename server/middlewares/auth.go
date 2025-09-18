package middlewares

import (
	"net/http"

	"api.link.henil.dev/internal/db/entities"
	"api.link.henil.dev/server/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkspaceContext struct {
	Workspace *entities.Workspace
}

const WorkspaceSecretHeaderName string = "x-workspace-secret"
const WorkspaceIdHeaderName string = "x-workspace-id"
const WorkspaceContextKey string = "workspace"

// check if workspace secret header is present
// if not present, forbid
// if present, check if valid
// if not valid, forbit
// if valid, set workspace in context and proceed
func WorkspaceSecretAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		workspaceSecret := c.GetHeader(WorkspaceSecretHeaderName)

		if workspaceSecret == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.NewErrorResponse("Workspace secret header is required"))
			return
		}

		workspace, err := entities.GetWorkspaceBySecret(c.Request.Context(), workspaceSecret)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.AbortWithStatusJSON(http.StatusForbidden, utils.NewErrorResponse("Invalid workspace secret"))
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.NewErrorResponse("Failed to authenticate workspace"))
			return
		}

		c.Set(WorkspaceContextKey, &WorkspaceContext{
			Workspace: workspace,
		})

		c.Next()
	}
}

// WorkspaceAuthMiddleware checks for both workspace ID and secret
// if either is missing, forbid
// if both are present, check if valid
// if not valid, forbid
// if valid, set workspace in context and proceed
func WorkspaceAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		workspaceId := c.GetHeader(WorkspaceIdHeaderName)
		if workspaceId == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.NewErrorResponse("Workspace ID header is required"))
			return
		}

		workspaceSecret := c.GetHeader(WorkspaceSecretHeaderName)
		if workspaceSecret == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.NewErrorResponse("Workspace secret header is required"))
			return
		}

		workspace, err := entities.GetWorkspaceByIDAndSecret(c.Request.Context(), workspaceId, workspaceSecret)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.AbortWithStatusJSON(http.StatusForbidden, utils.NewErrorResponse("Invalid workspace ID or secret"))
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.NewErrorResponse("Failed to authenticate workspace"))
			return
		}

		c.Set(WorkspaceContextKey, &WorkspaceContext{
			Workspace: workspace,
		})

		c.Next()
	}
}

func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
