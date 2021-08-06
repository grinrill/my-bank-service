package api

import (
	"net/http"
	"golang.org/x/xerrors"

	"github.com/gin-gonic/gin"
)

// GetAccountCurrencyRate validate query parameters
// and if all ok, call service.GetAccountCurrencyRate
func (h *Handler) GetAccountCurrencyRate(c *gin.Context) {
	cur := c.Query("cur")
	if cur=="" {
		c.JSON(http.StatusBadRequest, errorResponse(xerrors.New("cur is empty")))
		return
	}

	rate, err = h.bankService.GetAccountCurrencyRate(cur)
	if err!=nil {
		c.JSON(http.StatusBadGateway, errorResponse(err))
		return
	}

	// When all ok:
	c.JSON(http.StatusOK, okResponse(rate))
}