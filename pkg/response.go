package pkg

import (
	"github.com/gofiber/fiber/v2"
)

type MetaData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type StandardResponse struct {
	MetaData MetaData    `json:"metaData"`
	Response interface{} `json:"response"`
}

func Success(ctx *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return ctx.Status(statusCode).JSON(StandardResponse{
		MetaData: MetaData{
			Code:    statusCode,
			Message: message,
		},
		Response: data,
	})
}

func Error(ctx *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return ctx.Status(statusCode).JSON(StandardResponse{
		MetaData: MetaData{
			Code:    statusCode,
			Message: message,
		},
		Response: data,
	})
}

func PaginatedSuccess(ctx *fiber.Ctx, code int, message string, data interface{}, total, page, limit int) error {
	response := BuildPagination(ctx, data, total, page, limit)

	return ctx.Status(code).JSON(fiber.Map{
		"metaData": fiber.Map{
			"code":    code,
			"message": message,
		},
		"response": response,
	})
}