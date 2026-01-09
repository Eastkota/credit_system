package handlers

import (
	"credit_system/auth_service/graph"
	"credit_system/auth_service/model"

	"context"
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4" 
)
 type PaymentHandler struct {
	Service *service.PaymentService // Inject Service
}
func NewPaymentHandler(service *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{Service: service}
}
func (ph *PaymentHandler) CreatePayment(c echo.Context) error {
	var req model.CreatePaymentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	resp, err := ph.Service.CreateOwnerPayment(c.Request().Context(), req)
	if err != nil {
		switch err {
		case service.ErrInvalidInput:
			return c.JSON(http.StatusBadRequest, map[string]any{"error":"store_id, customer_id, amount are required"})
		case service.ErrOverPayment:
			return c.JSON(http.StatusBadRequest, map[string]any{"error":"payment exceeds outstanding balance"})
		case service.ErrBalanceNotFound:
			return c.JSON(http.StatusNotFound, map[string]string{"error": "customer balance not found"})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		}
	}
	return c.JSON(http.StatusCreated, resp)
}