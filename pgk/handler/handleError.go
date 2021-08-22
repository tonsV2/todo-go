package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tonsV2/todo-go/pgk/apperrors"
)

func HandleError(c *gin.Context, err error) {
	c.JSON(apperrors.ToHttpStatusCode(err), err.Error())
	c.Abort()
}
