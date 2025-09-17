package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBodyJSON[T any](c *gin.Context) (*T, bool) {
	var res *T
	err := c.ShouldBindJSON(&res)

	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse("Invalid request payload"))
		return nil, false
	}

	return res, true
}
