package handlers

import (
	"net/http"
	"time"

	"api.link.henil.dev/internal/db/entities"
	"api.link.henil.dev/server/middlewares"
	"api.link.henil.dev/server/utils"
	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */

type CreateApiKeyRequest struct {
	Label       string                      `json:"label" binding:"required"`
	Permissions []entities.ApiKeyPermission `json:"permissions" binding:"required"`
}

type CreateApiKeyResponse struct {
	ID  string `json:"id"`
	Key string `json:"key"`

	Label       string                      `json:"label"`
	Permissions []entities.ApiKeyPermission `json:"permissions"`
}

// CreateApiKey godoc
// @Summary Create a new API key
// @Description Create new API key for a workspace
// @ID create-api-key
// @Tags API Keys
// @Accept json
// @Produce json
// @Param x-workspace-id header string true "Workspace ID"
// @Param x-workspace-secret header string true "Workspace Secret"
// @Param api_key body CreateApiKeyRequest true "API Key details"
// @Success 201 {object} utils.Response{data=CreateApiKeyResponse} "API Key created successfully"
// @Failure 400 {object} utils.Response "Invalid request payload"
// @Failure 401 {object} utils.Response "Unauthorized"
// @Failure 500 {object} utils.Response "Failed to create API key"
// @Router /v1/api-keys [post]
func CreateApiKey(c *gin.Context) {

	body, ok := utils.GetBodyJSON[CreateApiKeyRequest](c)

	if !ok {
		return
	}

	workspaceCtx := c.MustGet(middlewares.WorkspaceContextKey).(*middlewares.WorkspaceContext)

	apiKey := entities.NewApiKey(body.Label, body.Permissions, workspaceCtx.Workspace.ID)

	id, err := apiKey.Insert(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Failed to create API key"))
		return
	}

	apiKey.ID = id

	c.JSON(http.StatusCreated, utils.NewSuccessResponse("API Key created successfully", CreateApiKeyResponse{
		ID:          apiKey.ID.Hex(),
		Key:         apiKey.Key,
		Label:       apiKey.Label,
		Permissions: apiKey.Permissions,
	}))
}

/* -------------------------------------------------------------------------- */

type GetAllApiKeysResponse struct {
	ApiKeys []ApiKeyDetails `json:"api_keys"`
}

type ApiKeyDetails struct {
	ID          string                      `json:"id"`
	Key         string                      `json:"key"`
	Label       string                      `json:"label"`
	Permissions []entities.ApiKeyPermission `json:"permissions"`
	CreatedAt   time.Time                   `json:"created_at"`
}

// GetAllApiKeys godoc
// @Summary Get all API keys
// @Description Retrieve all API keys for a workspace
// @ID get-all-api-keys
// @Tags API Keys
// @Accept json
// @Produce json
// @Param x-workspace-id header string true "Workspace ID"
// @Param x-workspace-secret header string true "Workspace Secret"
// @Success 200 {object} utils.Response{data=GetAllApiKeysResponse} "List of API keys"
// @Failure 401 {object} utils.Response "Unauthorized"
// @Failure 500 {object} utils.Response "Failed to retrieve API keys"
// @Router /v1/api-keys [get]
func GetAllApiKeys(c *gin.Context) {

	workspaceCtx := c.MustGet(middlewares.WorkspaceContextKey).(*middlewares.WorkspaceContext)

	apiKeys, err := entities.GetAllApiKeysByWorkspaceID(c.Request.Context(), workspaceCtx.Workspace.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Failed to retrieve API keys"))
		return
	}

	apiKeyDetails := make([]ApiKeyDetails, len(apiKeys))

	for i := range apiKeys {
		apiKeyDetails[i] = ApiKeyDetails{
			ID:          apiKeys[i].ID.Hex(),
			Key:         apiKeys[i].Key,
			Label:       apiKeys[i].Label,
			Permissions: apiKeys[i].Permissions,
			CreatedAt:   apiKeys[i].CreatedAt,
		}
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse("List of API keys", GetAllApiKeysResponse{
		ApiKeys: apiKeyDetails,
	}))
}
