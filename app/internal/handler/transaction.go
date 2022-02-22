package handler

import (
	"app/internal/model"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TransactionService service interface
type TransactionService interface {
	Transfer(ctx context.Context, transfer model.Transfer) (string, error)
}

// Transaction ...
type Transaction struct {
	service TransactionService
}

// NewTransaction ...
func NewTransaction(service TransactionService) *Transaction {
	return &Transaction{service}
}

// RegisterRoutes ...
func (t *Transaction) RegisterRoutes(router *echo.Echo) {
	router.POST("/transactions", t.Create)
}

// Create ...
func (t *Transaction) Create(c echo.Context) error {
	var req model.Transfer
	if err := c.Bind(&req); err != nil {
		log.Printf("binding failed, err: %v\n", err)
		return err
	}

	txID, err := t.service.Transfer(c.Request().Context(), req)
	if err != nil {
		log.Printf("service transfer failed, err: %v\n", err)
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{"tx_id": txID})
}
