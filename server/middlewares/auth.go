package middlewares

import "github.com/gin-gonic/gin"

func WorkspaceAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
