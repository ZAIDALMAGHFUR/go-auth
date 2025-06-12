package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/username/go-app/internal/auth/delivery/http/controller"
	authRepo "github.com/username/go-app/internal/auth/repository/pgsql"
	authRoutes "github.com/username/go-app/internal/auth/routes"
	authService "github.com/username/go-app/internal/auth/service"
)

func Init(app *fiber.App) {

	userRepo := authRepo.NewUserRepository()
	authService := authService.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)

	api := app.Group("/api/v1/services/mec@team")

	authRoutes.RegisterAuthRoutes(api, authController)
}
