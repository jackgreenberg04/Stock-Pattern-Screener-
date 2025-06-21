package internal

import "math"

func BullishEngulfing(c []Candle) bool {
	if len(c) < 2 {
		return false
	}
	prev := c[len(c)-2]
	curr := c[len(c)-1]
	return prev.Close < prev.Open && curr.Close > curr.Open && curr.Open < prev.Close && curr.Close > prev.Open
}

func BearishEngulfing(c []Candle) bool {
	if len(c) < 2 {
		return false
	}
	prev := c[len(c)-2]
	curr := c[len(c)-1]
	return prev.Close > prev.Open && curr.Close < curr.Open && curr.Open > prev.Close && curr.Close < prev.Open
}

func Hammer(c Candle) bool {
	body := math.Abs(c.Close - c.Open)
	candle := c.High - c.Low
	lowerShadow := c.Open - c.Low
	return body/candle <= 0.3 && lowerShadow/candle >= 0.5 && c.Close > c.Open
}

func InvertedHammer(c Candle) bool {
	body := math.Abs(c.Close - c.Open)
	candle := c.High - c.Low
	upper := c.High - c.Close
	return body/candle <= 0.3 && upper/candle >= 0.5 && c.Close > c.Open
}

func Doji(c Candle) bool {
	candle := c.High - c.Low
	if candle == 0 {
		return false
	}
	return math.Abs(c.Close-c.Open)/candle <= 0.1
}

func SMA(c []Candle, period int) []float64 {
	if len(c) < period {
		return nil
	}
	out := make([]float64, len(c)-period+1)
	for i := range out {
		sum := 0.0
		for j := 0; j < period; j++ {
			sum += c[i+j].Close
		}
		out[i] = sum / float64(period)
	}
	return out
}

func GoldenCross(c []Candle) bool {
	if len(c) < 200 {
		return false
	}
	sma50 := SMA(c, 50)
	sma200 := SMA(c, 200)
	idx := len(sma200) - 1
	if idx <= 0 {
		return false
	}
	return sma50[idx] > sma200[idx] && sma50[idx-1] <= sma200[idx-1]
}

func DeathCross(c []Candle) bool {
	if len(c) < 200 {
		return false
	}
	sma50 := SMA(c, 50)
	sma200 := SMA(c, 200)
	idx := len(sma200) - 1
	if idx <= 0 {
		return false
	}
	return sma50[idx] < sma200[idx] && sma50[idx-1] >= sma200[idx-1]
}

func RSI(c []Candle, period int) float64 {
	if len(c) <= period {
		return 50
	}
	gains, losses := 0.0, 0.0
	for i := len(c) - period; i < len(c)-1; i++ {
		change := c[i+1].Close - c[i].Close
		if change > 0 {
			gains += change
		} else {
			losses -= change
		}
	}
	if losses == 0 {
		return 100
	}
	rs := gains / losses
	return 100 - 100/(1+rs)
}
