package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tonsV2/todo-go/pgk/handler"
	"net/http"
)

func ProvideHandler(userService Service) Handler {
	return Handler{
		userService,
	}
}

type Handler struct {
	userService Service
}

func (h *Handler) Signup(c *gin.Context) {
	type signupRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,gte=16,lte=128"`
	}

	var request signupRequest

	if ok := handler.DataBinder(c, &request); !ok {
		return
	}

	user, err := h.userService.Signup(request.Email, request.Password)

	if err != nil {
		handler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h Handler) Me(c *gin.Context) {
	user, err := handler.GetUserFromContext(c)
	if err != nil {
		handler.HandleError(c, err)
		return
	}

	userWithGroups, err := h.userService.FindById(user.ID)
	if err != nil {
		handler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, userWithGroups)
}
