package http

import (
	"codelabs-backend-fiber/internal/user/domain"
	"codelabs-backend-fiber/internal/user/dto"
	customError "codelabs-backend-fiber/pkg/error"
	"codelabs-backend-fiber/pkg/response"
	"codelabs-backend-fiber/pkg/utils"
	"codelabs-backend-fiber/pkg/validator"
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase domain.UserUsecase
}

func NewUserHandler(uc domain.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: uc,
	}
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.usecase.GetAll()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Get List Users Success", users)
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := h.usecase.GetByID(uint(id))
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "User not found")
	}
	return response.Success(c, fiber.StatusOK, "Success Get Detail User", user)
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	if errors := validator.ValidateStruct(req); errors != nil {
		return response.Error(c, fiber.StatusBadRequest, fmt.Sprintf("validation_errors: %v", errors))
	}

	user := &domain.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
		Role:     domain.Role(req.Role),
	}

	err := h.usecase.Create(user)
	if err != nil {
		switch {
		case errors.Is(err, customError.ErrEmailAlreadyExists):
			return response.Error(c, fiber.StatusBadRequest, "email already exists")
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Internal Server Error")
		}
	}

	return response.Success(c, fiber.StatusCreated, "User created", map[string]any{})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request")
	}

	user, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "Invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email, string(user.Role))
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to generate token")
	}

	return response.Success(c, fiber.StatusOK, "Login success", fiber.Map{
		"token": token,
	})
}
