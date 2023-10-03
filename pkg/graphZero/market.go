package graphZero

type Market struct { // Market - биржа
	Name string // Name - название биржи
}

func NewMarket(
	name string,
) *Market {
	return &Market{
		Name: name,
	}
}

type OrderBook struct { // Книга ордеров (OrderBook) - два списка лотов - на покупку и на продажу
	Bids []*Lot `json:"bids"`
	Asks []*Lot `json:"asks"`
}

type Lot struct { // Lot - биржевой лот
	Price  float64 `json:"price"` // Price - Цена
	PriceQ float64 `json:"priceQ"`
	Amount float64 `json:"amount"` // Amount - Объём
}
