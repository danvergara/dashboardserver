package controllers

import (
	"net/http"
	"net/url"
	"time"

	"github.com/danvergara/dashservergo/currencyexchange"
	"github.com/gin-gonic/gin"
)

// CurrrencyExchange returns the currency exchange between the dollar and the mexican peso
func CurrrencyExchange(c *gin.Context) {
	currencyClient := currencyexchange.Client{}
	params := url.Values{}
	params.Add("base", "USD")
	params.Add("symbols", "MXN")

	resp, err := currencyClient.GetLatestCurrencyExchange(params)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"currency-exchange": resp})
}

// HistoricalCurrencyRates returns the historical currency rates given start and end dates
func HistoricalCurrencyRates(c *gin.Context) {
	currencyClient := currencyexchange.Client{}
	now := time.Now()
	dateLayout := "2006-01-02"

	params := url.Values{}
	params.Add("base", "USD")
	params.Add("symbols", "MXN")
	params.Add("end_at", now.Format(dateLayout))
	params.Add("start_at", now.AddDate(0, 0, -20).Format(dateLayout))
	resp, err := currencyClient.GetHistoricalCurrencyRate(params)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"historical-currency-rates": resp})
}
