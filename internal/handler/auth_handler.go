package handler

import (
	"net/http"
	"reast-api/internal/models"
	"reast-api/internal/service"
	"reast-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest

	// Validation mapping
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	if err := h.service.Register(req); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to register user", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User registered successfully", nil)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	token, err := h.service.Login(req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Login failed", err.Error())
		return
	}

	// Return token on success
	response.Success(c, http.StatusOK, "Login successful", gin.H{"token": token})
}
