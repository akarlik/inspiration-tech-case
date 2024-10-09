package handlers

import (
	"inspiration-tech-case/internal/models"
	"inspiration-tech-case/internal/services"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	TransactionService services.TransactionService
}

// HandleTransaction func handle new transaction.
// @Description handle new transaction.
// @Summary handle new transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param request body models.TransactionRequest true "Request Body"
// @Success 200 {object} models.GeneralResponse
// @Router /api/transactions [post]
func (handler *TransactionHandler) HandleTransaction(c *fiber.Ctx) error {
	var request models.TransactionRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
			Code:    400,
			Message: "Invalid request",
			Data:    nil,
		})
	}
	err := handler.TransactionService.HandleTransaction(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
			Code:    400,
			Message: "Invalid request",
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    true,
	})
}
