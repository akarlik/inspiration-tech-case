package handlers

import (
	"inspiration-tech-case/internal/models"
	"inspiration-tech-case/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	AccountService services.AccountService
}

// GetAccountByID func gets one exists account.
// @Description Get one exists account.
// @Summary get one exists account
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Account Id"
// @Success 200 {object} models.GeneralResponse
// @Router /api/accounts/{id} [get]
func (handler *AccountHandler) GetAccountByID(c *fiber.Ctx) error {
	id := c.Params("id")
	accountId, err := strconv.ParseUint(id, 10, 32)

	account, err := handler.AccountService.GetAccountByID(int64(accountId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.GeneralResponse{
			Code:    404,
			Message: "Not Found",
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    account,
	})

}
