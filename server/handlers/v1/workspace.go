package handlers

import (
	"net/http"

	"api.link.henil.dev/internal/db/entities"
	"api.link.henil.dev/server/utils"
	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */

type CreateWorkspaceRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateWorkspaceResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

var ResponseFailedToCreateWorkspace = utils.NewErrorResponse("Failed to create workspace")

// CreateWorkspace godoc
// @Summary Create a new workspace
// @Description Create new workspace using given name
// @ID create-workspace
// @Tags Workspaces
// @Accept json
// @Produce json
// @Param workspace body CreateWorkspaceRequest true "Workspace Name"
// @Success 200 {object} utils.Response{data=CreateWorkspaceResponse} "Workspace created successfully"
// @Failure 400 {object} utils.Response "Invalid request payload"
// @Router /v1/workspaces [post]
func CreateWorkspace(c *gin.Context) {
	body, ok := utils.GetBodyJSON[CreateWorkspaceRequest](c)
	if !ok {
		return
	}

	workspace := entities.NewWorkspaceByName(body.Name)

	id, err := workspace.Insert(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseFailedToCreateWorkspace)
		return
	}

	workspace.ID = id

	c.JSON(http.StatusOK, utils.NewSuccessResponse("Workspace created successfully", CreateWorkspaceResponse{
		ID:     workspace.ID.Hex(),
		Name:   workspace.Name,
		Secret: workspace.Secret,
	}))
}

/* -------------------------------------------------------------------------- */

// GetWorkspaceByID godoc
// @Summary Get workspace by ID
// @Description Retrieve workspace details using its ID
// @ID get-workspace-by-id
// @Tags Workspaces
// @Accept json
// @Produce json
// @Param id path string true "Workspace ID"
// @Param x-workspace-secret header string true "Workspace Secret"
// @Router /v1/workspaces/{id} [get]
func GetWorkspaceByID(c *gin.Context) {

}

/* -------------------------------------------------------------------------- */

// GetWorkspaceStatsByID godoc
// @Summary Get workspace stats by ID
// @Description Retrieve workspace statistics using its ID
// @ID get-workspace-stats-by-id
// @Tags Workspaces
// @Accept json
// @Produce json
// @Param id path string true "Workspace ID"
// @Param x-workspace-secret header string true "Workspace Secret"
// @Router /v1/workspaces/{id}/stats [get]
func GetWorkspaceStatsByID(c *gin.Context) {

}
