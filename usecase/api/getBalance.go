package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBalance validate query parameters
// if cur is empty, use service.GetCurrency instead
// and if all ok, call service.GetBalance
func (h *Handler) GetBalance(c *gin.Context) {
	cur := c.Query("cur")

	if cur=="" {
		cur, err := h.bankService.GetCurrency()
		if err!=nil {
			c.JSON(http.StatusBadGateway, errorResponse(err))
			return
		}
	}

	balance, err = h.bankService.GetBalance(cur)
	if err!=nil {
		c.JSON(http.StatusBadGateway, errorResponse(err))
		return
	}

	// When all ok:
	c.JSON(http.StatusOK, okResponse(balance))
}