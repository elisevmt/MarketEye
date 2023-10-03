package binanceAPI

import "time"

const (
	delayTickerPrice  = 50 * time.Millisecond
	delayTicker24hr   = 50 * time.Millisecond
	delayExchangeInfo = 50 * time.Millisecond
	delayOrderBook    = 50 * time.Millisecond
	delayTradeFee     = 200 * time.Millisecond
	delayServerTIme   = 50 * time.Millisecond
	delayOrder        = 50 * time.Millisecond
)
