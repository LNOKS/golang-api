package user

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService UserService
}

func NewMemberHandler(memberService UserService) *UserHandler {
	return &UserHandler{memberService}
}

func (h *UserHandler) GetAll(c echo.Context) error {
	searchTerm := c.QueryParam("searchTerm")
	users, err := h.userService.GetAll(searchTerm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response := make([]UserResponse, 0)
	for _, user := range users {
		response = append(response, UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) Create(c echo.Context) error {
	var member User
	err := c.Bind(&member)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = h.userService.Create(member)
	if err != nil {
		log.Errorf("Error creating member: %s", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, member)
}

func (h *UserHandler) Update(c echo.Context) error {
	var user User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = h.userService.Update(user)
	if err != nil {
		log.Errorf("Error updating user with id %d: %s", user.ID, err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	err := h.userService.Delete(idInt)
	if err != nil {
		log.Errorf("Error deleting user with id %d: %s", idInt, err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *UserHandler) DeleteBatch(c echo.Context) error {
	var ids []int
	err := c.Bind(&ids)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = h.userService.DeleteBatch(ids)
	if err != nil {
		log.Errorf("Error deleting users with ids %v: %s", ids, err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
