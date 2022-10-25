package handler

import (
	"net/http"
	"time"

	"github.com/AndreyDodonov/machclub_back/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input);
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

	logrus.Infof("user %s with email: %s was created :)", input.UserName, input.Email)

}

type signInInput struct {
	Email    string `json:"email" binding: "required,email"`
	Password string `json:"password" binding: "required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password);
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

	logrus.Infof("user %s is login at %s", input.Email, time.Now().Format("2006-01-02 15:04:05"))
}