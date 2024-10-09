package api

import (
	"inspiration-tech-case/configuration"
	"inspiration-tech-case/internal/api/handlers"
	"inspiration-tech-case/internal/repositories"
	"inspiration-tech-case/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

func SetupRoutes(app *fiber.App, db *cache.Cache, config configuration.Config) {
	accountRepo := repositories.NewAccountRepository(db, config)
	accountRepo.Seed()
	accountService := services.NewAccountService(&accountRepo)
	accountHandler := &handlers.AccountHandler{
		AccountService: accountService,
	}

	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(&transactionRepo)
	transactionHandler := &handlers.TransactionHandler{
		TransactionService: transactionService,
	}

	api := app.Group("/api")
	api.Get("/accounts/:id", accountHandler.GetAccountByID)
	api.Post("/transactions", transactionHandler.HandleTransaction)
}
