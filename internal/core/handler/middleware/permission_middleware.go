package middleware

import (
	"github.com/gin-gonic/gin"
	"goProject/internal"
	"goProject/internal/core/model"
	"goProject/pkg/jwt"
	"net/http"
)

func RoleAuth(required model.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusForbidden, internal.ErrUserNotFound)
			c.Abort()
			return
		}
		userToken, ok := user.(*jwt.UserToken)
		if !ok {
			c.JSON(http.StatusForbidden, internal.ErrUserNotFound)
			c.Abort()
			return
		}

		if userToken.Role != required {
			c.JSON(http.StatusForbidden, gin.H{"error": "you are not authorized to access this resource"})
			c.Abort()
			return
		}
		c.Next()
	}

}
