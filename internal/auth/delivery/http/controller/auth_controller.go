package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/username/go-app/internal/auth/delivery/http/request"
	"github.com/username/go-app/internal/auth/delivery/http/response"
	"github.com/username/go-app/internal/auth/service"
	"github.com/username/go-app/pkg"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	var req request.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return pkg.Error(ctx, fiber.StatusBadRequest, "Cannot parse JSON", nil)
	}

	user, err := c.authService.Register(req.Name, req.Email, req.Password)
	if err != nil {
		return pkg.Error(ctx, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return pkg.Success(ctx, fiber.StatusOK, "Registration success", response.FromDomain(user))
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var req request.LoginRequest
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
