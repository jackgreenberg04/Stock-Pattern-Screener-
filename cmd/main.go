package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"Stock-Pattern-Screener-/internal"
)

func main() {
	cfg, err := internal.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	logFile, err := os.OpenFile("patterns.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("log file: %v", err)
	}
	defer logFile.Close()
	logger := log.New(logFile, "", log.LstdFlags)

	results := make(map[string]internal.Result)
	candleStore := make(map[string][]internal.Candle)
	fetchAndAnalyze := func() {
		for _, t := range cfg.Tickers {
			candles, err := internal.FetchHistory(t, cfg.Interval, cfg.Range)
			if err != nil {
				log.Println("fetch", t, err)
				continue
			}
			candleStore[t] = candles
			matched := []string{}
			if internal.BullishEngulfing(candles) {
				matched = append(matched, "Bullish Engulfing")
			}
			if internal.BearishEngulfing(candles) {
				matched = append(matched, "Bearish Engulfing")
			}
			last := candles[len(candles)-1]
			if internal.Hammer(last) {
				matched = append(matched, "Hammer")
			}
			if internal.InvertedHammer(last) {
				matched = append(matched, "Inverted Hammer")
			}
			if internal.Doji(last) {
				matched = append(matched, "Doji")
			}
			if internal.GoldenCross(candles) {
				matched = append(matched, "Golden Cross")
			}
			if internal.DeathCross(candles) {
				matched = append(matched, "Death Cross")
			}
			rsi := internal.RSI(candles, 14)
			if rsi > 70 {
				matched = append(matched, "RSI Overbought")
			} else if rsi < 30 {
				matched = append(matched, "RSI Oversold")
			}
			res := internal.Result{Ticker: t, Patterns: matched, Time: time.Now()}
			results[t] = res
			if len(matched) > 0 {
				logger.Println(t, matched)
			}
		}
	}

	fetchAndAnalyze()
	ticker := time.NewTicker(time.Duration(cfg.RefreshMinutes) * time.Minute)
	go func() {
		for range ticker.C {
			fetchAndAnalyze()
		}
	}()

	getResults := func() []internal.Result {
		out := make([]internal.Result, 0, len(results))
		for _, v := range results {
			out = append(out, v)
		}
		return out
	}

	getCandles := func(t string) []internal.Candle {
		return candleStore[t]
	}

	gin.SetMode(gin.ReleaseMode)
	srv := internal.NewServer(time.Now(), cfg.APIKey, getResults, getCandles)
	log.Fatal(srv.Run(":8080"))
}
