package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mattevans/dinero"
)

// CurrrencyExchange returns the currency exchange between the dollar and the mexican peso
func CurrrencyExchange(c *gin.Context) {
	currencyClient := dinero.NewClient(
		os.Getenv("OPEN_EXCHANGE_APP_ID"),
		"USD",
		20*time.Minute,
	)

	rsp, err := currencyClient.Rates.Get("MXN")

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"dolar-peso": rsp})
}
