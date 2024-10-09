package tests

import (
	"bytes"
	"encoding/json"
	"inspiration-tech-case/internal/api/handlers"
	"inspiration-tech-case/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionService struct {
	mock.Mock
}

func (m *MockTransactionService) HandleTransaction(request models.TransactionRequest) error {
	args := m.Called(request)
	return args.Error(1)
}

func TestHandlerHandleTransaction_BadRequest(t *testing.T) {
	app := fiber.New()

	mockService := new(MockTransactionService)
	transactionRequest := models.TransactionRequest{TransactionId: "1"}
	mockService.On("HandleTransaction", transactionRequest).Return(nil)

	handler := &handlers.TransactionHandler{TransactionService: mockService}
	api := app.Group("/api")
	api.Post("/transactions", handler.HandleTransaction)

	body, err := json.Marshal(transactionRequest)
	if err != nil {
		t.Errorf("Cannot parse body: %v", err)
		t.Fail()
	}
	req := httptest.NewRequest("POST", "/api/transactions", bytes.NewReader(body))
	resp, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

}


