# Stock Pattern Screener

Stock Pattern Screener periodically downloads historical OHLCV data for a set of configured tickers and checks for common bullish and bearish technical patterns. A minimal REST API exposes the detected patterns and service status.

## Features

- Fetch daily stock prices from Yahoo Finance
- Detect candlestick patterns (bullish/bearish engulfing, hammer, inverted hammer, doji)
- Check moving average crossovers and RSI
- Schedule automatic updates
- Log pattern matches to `patterns.log`
- REST endpoints for status, screening results and optional chart data

## Tech Stack

- **Go** 1.20
- **Gin** web framework
- Yahoo Finance HTTP API

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install)
- [Git](https://git-scm.com/)

### Clone the repository

```bash
git clone <repo-url>
cd Stock-Pattern-Screener-
```

### Install dependencies

```bash
go mod tidy
```

### Run

```bash
go run ./cmd
```

The service listens on `:8080`.

### Configuration

The application reads `config.json` in the working directory.

```
{
  "tickers": ["AAPL", "MSFT", "TSLA"],
  "interval": "1d",
  "range": "1mo",
  "refresh_minutes": 60,
  "api_key": "changeme"
}
```

- **tickers** – list of symbols to track
- **interval** – data interval (e.g. `1d`)
- **range** – data range passed to Yahoo Finance API
- **refresh_minutes** – how often to refresh prices
- **api_key** – key required when calling protected endpoints

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/status` | Service uptime |
| `GET` | `/screen?key=APIKEY` | JSON list of pattern matches |
| `GET` | `/chart/:ticker?key=APIKEY` | Returns raw candle data |

## Pattern Descriptions

- **Bullish Engulfing** – current candle completely engulfs previous bearish candle.
- **Bearish Engulfing** – current bearish candle engulfs previous bullish candle.
- **Hammer** – small body with long lower shadow after a downtrend.
- **Inverted Hammer** – small body with long upper shadow.
- **Doji** – open and close are nearly equal.
- **Golden Cross** – 50‑day SMA crosses above 200‑day SMA.
- **Death Cross** – 50‑day SMA crosses below 200‑day SMA.
- **RSI Overbought/Oversold** – RSI > 70 or < 30.

## License

MIT

## Author

Your Name
