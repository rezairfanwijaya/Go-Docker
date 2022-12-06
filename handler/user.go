package handler

import (
	"go-docker/helper"
	"go-docker/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.IUserService
}

func NewHandlerUser(userService user.IUserService) *UserHandler {
	return &UserHandler{userService}
}

// implementasi
func (h *UserHandler) Login(c *gin.Context) {
	var input user.UserInputLogin

	// binding
	if err := c.BindJSON(&input); err != nil {
		errBinding := helper.ErrorBindingFormatter(err)
		response := helper.ResponseAPIFormatter(
			"gagal",
			"gagal binding",
			http.StatusInternalServerError,
			errBinding,
		)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// panggil service
	user, err := h.userService.Save(input)
	if err != nil {
		response := helper.ResponseAPIFormatter(
			"gagal",
			"gagal login",
			http.StatusBadRequest,
			err.Error(),
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseAPIFormatter(
		"sukses",
		"sukses login",
		http.StatusOK,
		helper.UserFormatter(user),
	)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetAllUser(c *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		response := helper.ResponseAPIFormatter(
			"gagal",
			"gagal mengambil data user",
			http.StatusInternalServerError,
			err,
		)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var userResponse interface{}

	userFormatted := helper.UsersFormatter(users)
	userResponse = userFormatted
	if len(users) == 0 {
		userResponse = users
	}

	response := helper.ResponseAPIFormatter(
		"sukses",
		"sukses login",
		http.StatusOK,
		userResponse,
	)
	c.JSON(http.StatusOK, response)

}
