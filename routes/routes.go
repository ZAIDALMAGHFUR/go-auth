package route

import (
	"github.com/gofiber/fiber/v2"

	authController "github.com/username/go-app/internal/auth/delivery/http/controller"
	authRepo "github.com/username/go-app/internal/auth/repository/pgsql"
	authRoutes "github.com/username/go-app/internal/auth/routes"
	authService "github.com/username/go-app/internal/auth/service"

	agamaController "github.com/username/go-app/internal/master/agama/delivery/http/controller"
	agamaRepo "github.com/username/go-app/internal/master/agama/repository/pgsql"
	agamaRoutes "github.com/username/go-app/internal/master/agama/routes"
	agamaService "github.com/username/go-app/internal/master/agama/service"

	"github.com/username/go-app/internal/middleware"
)

func Init(app *fiber.App) {
	app.Use(middleware.Logger())
	app.Use(middleware.Recovery())

	userRepo := authRepo.NewUserRepository()
	authSvc := authService.NewAuthService(userRepo)
	authCtrl := authController.NewAuthController(authSvc)

	agamaRepo := agamaRepo.NewAgamaRepository()
	agamaSvc := agamaService.NewAgamaService(agamaRepo)
	agamaCtrl := agamaController.NewAgamaController(agamaSvc)

	publicAPI := app.Group("/api/v1/services/mec@team")
	authRoutes.RegisterAuthRoutes(publicAPI, authCtrl)

	protectedAPI := app.Group("/api/v1/services/mec@team", middleware.AuthMiddleware())
	agamaRoutes.RegisterAgamaRoutes(protectedAPI, agamaCtrl)
}
