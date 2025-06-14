package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zaidalmaghfur/go-app/internal/master/agama/delivery/http/request"
	"github.com/zaidalmaghfur/go-app/internal/master/agama/delivery/http/response"
	"github.com/zaidalmaghfur/go-app/internal/master/agama/service"
	"github.com/zaidalmaghfur/go-app/pkg"
)

type AgamaController struct {
	agamaService service.AgamaService
}

func NewAgamaController(service service.AgamaService) *AgamaController {
	return &AgamaController{agamaService: service}
}

func (c *AgamaController) Create(ctx *fiber.Ctx) error {
	var req request.AgamaRequest
	if err := ctx.BodyParser(&req); err != nil {
		return pkg.Error(ctx, fiber.StatusBadRequest, "Cannot parse JSON", nil)
	}
	if ok, errs := pkg.ValidateStruct(req); !ok {
		return pkg.Error(ctx, fiber.StatusBadRequest, "Validation failed", errs)
	}
	agama, err := c.agamaService.Create(req.Name)
	if err != nil {
		return pkg.Error(ctx, fiber.StatusInternalServerError, err.Error(), nil)
	}
	return pkg.Success(ctx, fiber.StatusCreated, "Created successfully", response.FromDomain(agama))
}

func (c *AgamaController) GetByID(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return pkg.Error(ctx, fiber.StatusBadRequest, "Invalid ID", nil)
	}
	agama, err := c.agamaService.GetByID(uint(id))
	if err != nil {
		return pkg.Error(ctx, fiber.StatusNotFound, err.Error(), nil)
	}
	return pkg.Success(ctx, fiber.StatusOK, "Retrieved successfully", response.FromDomain(agama))
}

func (c *AgamaController) Update(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return pkg.Error(ctx, fiber.StatusBadRequest, "Invalid ID", nil)
	}
	var req request.AgamaRequest
	if err := ctx.BodyParser(&req); err != nil {
		return pkg.Error(ctx, fiber.StatusBadRequest, "Cannot parse JSON", nil)
	}
	if ok, errs := pkg.ValidateStruct(req); !ok {
		return pkg.Error(ctx, fiber.StatusBadRequest, "Validation failed", errs)
	}
	agama, err := c.agamaService.Update(uint(id), req.Name)
	if err != nil {
		return pkg.Error(ctx, fiber.StatusInternalServerError, err.Error(), nil)
	}
	return pkg.Success(ctx, fiber.StatusOK, "Updated successfully", response.FromDomain(agama))
}

func (c *AgamaController) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return pkg.Error(ctx, fiber.StatusBadRequest, "Invalid ID", nil)
	}
	if err := c.agamaService.Delete(uint(id)); err != nil {
		return pkg.Error(ctx, fiber.StatusInternalServerError, err.Error(), nil)
	}
	return pkg.Success(ctx, fiber.StatusOK, "Deleted successfully", nil)
}

func (c *AgamaController) GetAll(ctx *fiber.Ctx) error {
	pageStr := ctx.Query("page", "1")
	limitStr := ctx.Query("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	agamas, total, err := c.agamaService.GetAllPaginated(offset, limit)
	if err != nil {
		return pkg.Error(ctx, fiber.StatusInternalServerError, err.Error(), nil)
	}

	data := response.FromDomainList(agamas)

	paginated := pkg.Build(ctx, data, total, page, limit)

	return pkg.Success(ctx, fiber.StatusOK, "Module retrieved successfully", paginated)
}