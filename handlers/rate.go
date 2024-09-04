package handlers

import (
	"github.com/gin-gonic/gin"
	"moneytracker-backend/coingate"
	"net/http"
	"time"
)

func Rate(c *gin.Context) {
	from := c.Query("from")
	if from == "" {
		c.String(http.StatusBadRequest, "from field is required")
		return
	}

	to := c.Query("to")
	if to == "" {
		c.String(http.StatusBadRequest, "to field is required")
		return
	}

	// TODO: from and to should be validated against a list of supported currencies
	rate, err := coingate.CurrencyRate(coingate.CurrencyRateOptions{
		From:     from,
		To:       to,
		CacheFor: 5 * time.Minute,
	})

	if err != nil {
		c.String(http.StatusInternalServerError, "error: %s", err.Error())
		return
	}

	c.String(http.StatusOK, rate)
}
