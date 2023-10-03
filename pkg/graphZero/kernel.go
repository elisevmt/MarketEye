package graphZero

type KernelPrice struct {
	Symbol       string  `json:"symbol"`
	Left         string  `json:"left"`
	Right        string  `json:"right"`
	Price        float64 `json:"price"`
	PriceQ       float64 `json:"priceQ"`
	AveragePrice float64 `json:"averagePrice"`
}

type KernelPriceList struct {
	Market string        `json:"market"`
	Data   []KernelPrice `json:"data"`
}

type KernelMarketPrice struct {
	Markets []KernelPriceList `json:"markets"`
}
