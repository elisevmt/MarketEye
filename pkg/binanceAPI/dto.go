package binanceAPI

type BatchOrderBookWSS struct {
	Stream string       `json:"stream"`
	Data   OrderBookWSS `json:"data"`
}

type OrderBookWSS struct {
	EventType     string     `json:"e"`
	EventTime     int64      `json:"E"`
	Symbol        string     `json:"s"`
	UpdateFirstID int64      `json:"U"`
	UpdateLastID  int64      `json:"u"`
	Bids          [][]string `json:"b"`
	Asks          [][]string `json:"a"`
}

type Ticker24hWSS struct {
	EventType            string `json:"e"`
	EventTime            int64  `json:"E"`
	Symbol               string `json:"s"`
	PriceChange          string `json:"p"`
	PriceChangePercent   string `json:"P"`
	WeightedAveragePrice string `json:"w"`
	X                    string `json:"x"`
	LastPrice            string `json:"c"`
	LastQuantity         string `json:"Q"`
	BestBidPrice         string `json:"b"`
	BestBidQuantity      string `json:"B"`
	BestAskPrice         string `json:"a"`
	BestAskQuantity      string `json:"A"`
	OpenPrice            string `json:"o"`
	H                    string `json:"h"`
	L                    string `json:"l"`
	V                    string `json:"v"`
	Q1                   string `json:"q"`
	O1                   int64  `json:"O"`
	C1                   int64  `json:"C"`
	F                    int    `json:"F"`
	L1                   int    `json:"L"`
	N                    int    `json:"n"`
}

type Ticker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type Ticker24hr struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstId            int    `json:"firstId"`
	LastId             int    `json:"lastId"`
	Count              int    `json:"count"`
}

type ExchangeInfo struct {
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	RateLimits      []struct {
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
		RateLimitType string `json:"rateLimitType"`
	} `json:"rateLimits"`
	ServerTime int64 `json:"serverTime"`
	Symbols    []struct {
		BaseAsset               string `json:"baseAsset"`
		BaseAssetPrecision      int    `json:"baseAssetPrecision"`
		BaseCommissionPrecision int    `json:"baseCommissionPrecision"`
		Filters                 []struct {
			FilterType          string      `json:"filterType"`
			MinPrice            string      `json:"minPrice,omitempty"`
			MaxPrice            string      `json:"maxPrice,omitempty"`
			TickSize            string      `json:"tickSize,omitempty"`
			MultiplierUp        string      `json:"multiplierUp,omitempty"`
			MultiplierDown      string      `json:"multiplierDown,omitempty"`
			AvgPriceMins        int         `json:"avgPriceMins,omitempty"`
			MinQty              string      `json:"minQty,omitempty"`
			MaxQty              string      `json:"maxQty,omitempty"`
			StepSize            string      `json:"stepSize,omitempty"`
			MinNotional         string      `json:"minNotional,omitempty"`
			ApplyToMarket       interface{} `json:"applyToMarket,omitempty"`
			MaxPosition         string      `json:"maxPosition,omitempty"`
			MaxNumIcebergOrders string      `json:"maxNumIcebergOrders,omitempty"`
			Limit               int         `json:"limit,omitempty"`
			MaxNumOrders        int         `json:"maxNumOrders,omitempty"`
			MaxNumAlgoOrders    int         `json:"maxNumAlgoOrders,omitempty"`
		} `json:"filters"`
		IcebergAllowed             interface{} `json:"icebergAllowed"`
		IsMarginTradingAllowed     interface{} `json:"isMarginTradingAllowed"`
		IsSpotTradingAllowed       interface{} `json:"isSpotTradingAllowed"`
		OcoAllowed                 interface{} `json:"ocoAllowed"`
		OrderTypes                 []string    `json:"orderTypes"`
		Permissions                []string    `json:"permissions"`
		QuoteAsset                 string      `json:"quoteAsset"`
		QuoteAssetPrecision        int         `json:"quoteAssetPrecision"`
		QuoteCommissionPrecision   int         `json:"quoteCommissionPrecision"`
		QuoteOrderQtyMarketAllowed interface{} `json:"quoteOrderQtyMarketAllowed"`
		QuotePrecision             int         `json:"quotePrecision"`
		Status                     string      `json:"status"`
		Symbol                     string      `json:"symbol"`
	} `json:"symbols"`
}

type SourceBook struct {
	LastUpdateId int64      `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

type TradeFee struct {
	Symbol          string `json:"symbol"`
	MakerCommission string `json:"makerCommission"`
	TakerCommission string `json:"takerCommission"`
}

type ServerTime struct {
	Time int64 `json:"serverTime"`
}

type NewOrderResponse struct {
	Symbol              string `json:"symbol"`
	OrderID             int64  `json:"orderId"`
	OrderListID         int    `json:"orderListId"`
	ClientOrderID       string `json:"clientOrderId"`
	TransactTime        int64  `json:"transactTime"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	Fills               []struct {
		Price           string `json:"price"`
		Qty             string `json:"qty"`
		Commission      string `json:"commission"`
		CommissionAsset string `json:"commissionAsset"`
		TradeID         int    `json:"tradeId"`
	} `json:"fills"`
}

type NewOrderResponseShort struct {
	OrigQty             float64 `json:"origQty"`
	CummulativeQuoteQty float64 `json:"cummulativeQuoteQty"`
	AVGPrice            float64 `json:"avgPrice"`
	TransactTime        int64   `json:"transactTime"`
	Status              string  `json:"status"`
}
