package api

import (
	"github.com/grinrill/my-bank-service/domain/bankService"
)

// NewHandler create new Handler
func NewHandler(bankService bankService.AccountInterface) *Handler {
	return &Handler{bankService}
}

// Handler is service
// with http gin handlers
type Handler struct {
	bankService bankService.AccountInterface
}