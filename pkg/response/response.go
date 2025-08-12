package response

import (
	"github.com/gofiber/fiber/v2"
)

type PaginationMeta struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"totalPages"`
}

func Success(c *fiber.Ctx, status int, message string, data any) error {
    return c.Status(status).JSON(fiber.Map{
        "code":   status,
        "status": "success",
        "message": message,
        "data":   data,
    })
}

func Paginated(ctx *fiber.Ctx, message string, data any, meta PaginationMeta) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"status":  "success",
		"message": message,
		"data":    data,
		"meta":    meta,
	})
}

func Error(ctx *fiber.Ctx, code int, message string) error {
	return ctx.Status(code).JSON(fiber.Map{
		"code":    code,
		"status":  "error",
		"message": message,
		"data":    nil,
	})
}
