package api

import (
	"strconv"
	"net/http"
	"golang.org/x/xerrors"

	"github.com/gin-gonic/gin"
)

// Withdraw validate query parameters
// and if all ok, call service.Withdraw
func (h *Handler) Withdraw(c *gin.Context) {
	sumString := c.Query("sum")
	if sumString=="" {
		c.JSON(http.StatusBadRequest, errorResponse(xerrors.New("sum is empty")))
		return
	}

	sum, err := strconv.ParseFloat(sumString, 64)
	if err!=nil {
		c.JSON(http.StatusBadRequest, errorResponse(xerrors.New("sum is not valid float")))
		return
	}

	// service operation:
	err = h.bankService.Withdraw(sum)
	if err!=nil {
		c.JSON(http.StatusBadGateway, errorResponse(err))
		return
	}

	// When all ok:
	c.JSON(http.StatusOK, okResponse(nil))
}