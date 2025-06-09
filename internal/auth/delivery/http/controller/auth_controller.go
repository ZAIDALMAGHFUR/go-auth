package controller

import (
	"github.com/username/go-app/internal/auth/service"
	"github.com/username/go-app/pkg"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	type request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req request
	if err := ctx.BodyParser(&req); err != nil {
		return pkg.Error(ctx, fiber.StatusBadRequest, "Cannot parse JSON", nil)
	}

	user, err := c.authService.Register(req.Name, req.Email, req.Password)
	if err != nil {
		return pkg.Error(ctx, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return pkg.Success(ctx, fiber.StatusOK, "Registration success", user)
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req request
	if err := ctx.BodyParser(&req); err != nil {
		return pkg.Error(ctx, fiber.StatusBadRequest, "Cannot parse JSON", nil)
	}

	token, err := c.authService.Login(req.Email, req.Password)
	if err != nil {
		return pkg.Error(ctx, fiber.StatusUnauthorized, "Invalid credentials", nil)
	}

	return pkg.Success(ctx, fiber.StatusOK, "Login success", fiber.Map{
		"token": token,
	})
}
