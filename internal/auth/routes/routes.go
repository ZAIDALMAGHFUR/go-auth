package routes

import (
	"github.com/username/go-app/internal/auth/delivery/http/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(router fiber.Router, authController *controller.AuthController) {
	authGroup := router.Group("/auth")
	authGroup.Post("/register", authController.Register)
	authGroup.Post("/login", authController.Login)
}
