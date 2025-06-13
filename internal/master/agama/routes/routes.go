package routes

import (
	"github.com/username/go-app/internal/master/agama/delivery/http/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterAgamaRoutes(router fiber.Router, agamaController *controller.AgamaController) {
	agamaGroup := router.Group("/master")
	agamaGroup.Post("/agama", agamaController.Create)
	agamaGroup.Get("/agama/:id", agamaController.GetByID)
	agamaGroup.Put("/agama/:id", agamaController.Update)
	agamaGroup.Delete("/agama/:id", agamaController.Delete)
	agamaGroup.Get("/agama", agamaController.GetAll)
}