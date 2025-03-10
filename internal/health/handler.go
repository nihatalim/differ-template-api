package health

import "github.com/gofiber/fiber/v2"

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) HandleHealthCheck(c *fiber.Ctx) error {
	result := h.service.CheckHealth()
	return c.JSON(result)
} 