package internal

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Ticker   string    `json:"ticker"`
	Patterns []string  `json:"patterns"`
	Time     time.Time `json:"time"`
}

func NewServer(start time.Time, apiKey string, results func() []Result, candles func(string) []Candle) *gin.Engine {
	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"uptime": time.Since(start).String()})
	})

	auth := func(c *gin.Context) {
		if c.Query("key") != apiKey {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}

	r.GET("/screen", auth, func(c *gin.Context) {
		c.JSON(http.StatusOK, results())
	})

	r.GET("/chart/:ticker", auth, func(c *gin.Context) {
		t := c.Param("ticker")
		c.JSON(http.StatusOK, candles(t))
	})

	return r
}
