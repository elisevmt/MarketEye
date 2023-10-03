package models

type FetchPricesParams struct {
	MarketList []*string `json:"market_list,omitempty"`
	Node       *string   `json:"node,omitempty"`
	Left       *string   `json:"left,omitempty"`
	Right      *string   `json:"right,omitempty"`
	Symbol     *string   `json:"symbol,omitempty"`
}

type FetchOrderBookParams struct {
	Market string `json:"market"`
	Symbol string `json:"symbol"`
	Depth  *int64 `json:"limit,omitempty"`
}

type FetchMarketListResponse struct {
	MarketList []string `json:"market_list"`
}
