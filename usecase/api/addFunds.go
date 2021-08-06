package api

import (
	"strconv"
	"net/http"
	"golang.org/x/xerrors"

	"github.com/gin-gonic/gin"
)

// AddFunds validate query parameters
// and if all ok, call service.AddFunds
func (h *Handler) AddFunds(c *gin.Context) {
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
	err = h.bankService.AddFunds(sum)
	if err!=nil {
		c.JSON(http.StatusBadGateway, errorResponse(err))
		return
	}

	// When all ok:
	c.JSON(http.StatusOK, okResponse(nil))
}