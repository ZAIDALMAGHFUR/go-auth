package main

import (
	"github.com/username/go-app/config"
	authController "github.com/username/go-app/internal/auth/delivery/http/controller"
	"github.com/username/go-app/internal/auth/domain"
	authRepo "github.com/username/go-app/internal/auth/repository/mysql"
	authRoutes "github.com/username/go-app/internal/auth/routes"
	authService "github.com/username/go-app/internal/auth/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDatabase()

	config.DB.AutoMigrate(&domain.User{})

	userRepo := authRepo.NewUserRepository()
	authService := authService.NewAuthService(userRepo)
	authController := authController.NewAuthController(authService)

	api := app.Group("/api/v1")
	authRoutes.RegisterAuthRoutes(api, authController)

	app.Listen(":3000")
}
