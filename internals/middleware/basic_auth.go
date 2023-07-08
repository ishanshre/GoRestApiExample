package middleware

import "github.com/gin-gonic/gin"

func Basic_Auth() gin.HandlerFunc {
	return gin.BasicAuth(
		gin.Accounts{
			"admin": "admin",
		},
	)
}