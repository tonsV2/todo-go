package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tonsV2/todo-go/pgk/apperrors"
	"github.com/tonsV2/todo-go/pgk/user"
)

func GetUserFromContext(c *gin.Context) (*user.User, error) {
	u, exists := c.Get("user")

	if !exists {
		message := fmt.Sprintf("Unable to extract user from request context for unknown reason: %+v", c)
		err := apperrors.NewInternal(message)
		return nil, err
	}

	return u.(*user.User), nil
}
