package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tonsV2/todo-go/pgk/apperrors"
)

func DataBinder(c *gin.Context, req interface{}) bool {
	if c.ContentType() != "application/json" && c.ContentType() != "multipart/form-data" {
		reason := fmt.Sprintf("%s only accepts Content-Type application/json", c.FullPath())

		err := apperrors.NewUnsupportedMediaType(reason)

		c.JSON(err.HttpStatusCode(), err)
		c.Abort()

		return false
	}

	if err := c.ShouldBind(req); err != nil {
		message := fmt.Sprintf("Error binding data: %+v\n", err)
		err := apperrors.NewBadRequest(message)
		HandleError(c, err)
		return false
	}

	return true
}
