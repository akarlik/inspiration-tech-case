package tests

import (
	"encoding/json"
	"errors"
	"inspiration-tech-case/internal/api/handlers"
	"inspiration-tech-case/internal/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAccountService struct {
	mock.Mock
}

func (m *MockAccountService) GetAccountByID(accountId int64) (models.AccountResponse, error) {
	args := m.Called(accountId)
	return args.Get(0).(models.AccountResponse), args.Error(1)
}

func TestHandlerGetAccountByID(t *testing.T) {
	app := fiber.New()

	mockService := new(MockAccountService)
	account := models.AccountResponse{AccountId: 1, Balance: 2.0}
	mockService.On("GetAccountByID", int64(1)).Return(account, nil)

	handler := &handlers.AccountHandler{AccountService: mockService}
	api := app.Group("/api")
	api.Get("/accounts/:id", handler.GetAccountByID)

	req := httptest.NewRequest("GET", "/api/accounts/1", nil)
	resp, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Cannot parse body: %v", err)
	}

	var generalResponse models.GeneralResponse
	if err := json.Unmarshal(responseBody, &generalResponse); err != nil {
		t.Errorf("Cannot unmarshal body, might returned wrong type of output: %v", err)
	}
	mockService.AssertCalled(t, "GetAccountByID", int64(1))
}

func TestHandlerGetAccountByID_NotFound(t *testing.T) {
	app := fiber.New()

	mockService := new(MockAccountService)
	mockService.On("GetAccountByID", int64(2)).Return(models.AccountResponse{}, errors.New("account not found"))

	handler := &handlers.AccountHandler{AccountService: mockService}
	api := app.Group("/api")
	api.Get("/accounts/:id", handler.GetAccountByID)

	req := httptest.NewRequest("GET", "/api/accounts/2", nil)
	resp, _ := app.Test(req)

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Cannot parse body: %v", err)
	}

	var generalResponse models.GeneralResponse
	if err := json.Unmarshal(responseBody, &generalResponse); err != nil {
		t.Errorf("Cannot unmarshal body, might returned wrong type of output: %v", err)
	}

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	mockService.AssertCalled(t, "GetAccountByID", int64(2))
}
