package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Candle struct {
	Timestamp int64
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
}

// FetchHistory retrieves historical data from Yahoo Finance v8 chart API
func FetchHistory(ticker, interval, rng string) ([]Candle, error) {
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?range=%s&interval=%s", ticker, rng, interval)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result struct {
		Chart struct {
			Result []struct {
				Timestamp  []int64 `json:"timestamp"`
				Indicators struct {
					Quote []struct {
						Open   []float64 `json:"open"`
						High   []float64 `json:"high"`
						Low    []float64 `json:"low"`
						Close  []float64 `json:"close"`
						Volume []float64 `json:"volume"`
					} `json:"quote"`
				} `json:"indicators"`
			} `json:"result"`
			Error interface{} `json:"error"`
		} `json:"chart"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if len(result.Chart.Result) == 0 {
		return nil, fmt.Errorf("no data")
	}
	r := result.Chart.Result[0]
	quotes := r.Indicators.Quote[0]
	candles := make([]Candle, len(r.Timestamp))
	for i := range r.Timestamp {
		candles[i] = Candle{
			Timestamp: r.Timestamp[i],
			Open:      quotes.Open[i],
			High:      quotes.High[i],
			Low:       quotes.Low[i],
			Close:     quotes.Close[i],
			Volume:    quotes.Volume[i],
		}
	}
	return candles, nil
}
