package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCurrency call service.GetCurrency
func (h *Handler) GetCurrency(c *gin.Context) {
	currency, err = h.bankService.GetCurrency()
	if err!=nil {
		c.JSON(http.StatusBadGateway, errorResponse(err))
		return
	}

	// When all ok:
	c.JSON(http.StatusOK, okResponse(currency))
}