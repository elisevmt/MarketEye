package graphZero

import "fmt"

type Branch struct {
	Market        *Market        `json:"market"`         // Market - Биржа
	Left          *Node          `json:"left"`           // Left - Первая валюта
	Right         *Node          `json:"right"`          // Right - Вторая валюта
	LRWeight      float64        `json:"lr_weight"`      // LRWeight - Цена перехода от первой валюты ко второй
	WeightSize    int64          `json:"weight_size"`    // WeightSize - Количество знаков после запятой в цене перехода
	ExCount       int64          `json:"count"`          // ExCount - Количество обменов за 24 часа по данной валюте
	ExSymbol      string         `json:"ex_symbol"`      // ExSymbol - исходный символ на бирже, типа не бывает USDTBTC, исходное - BTCUSDT, он тут и хранится
	Bids          []*Lot         `json:"bids"`           // Bids - Список лотов на продажу
	Asks          []*Lot         `json:"asks"`           // Asks - Список лотов на покупку
	LeftWidth     float64        `json:"left_width"`     // LeftWidth - ширина канала с левой стороны
	RightWidth    float64        `json:"right_width"`    // RightWidth - ширина канала с правой стороны
	LRWeightQ     float64        `json:"lr_weight_q"`    // LRWeightQ - умная цена, пока не используется, но вдруг пригодится, оставим
	AveragePrice  float64        `json:"average_price"`  // AveragePrice - средняя цена биржи
	BranchFilters *BranchFilters `json:"branch_filters"` // BranchFilters - фильтры продажи, округление левого и правого валютного узла, а так же комиссия при продаже/покупке
}

type BranchFilters struct {
	LeftPrecision  int64   `json:"left_precision"`  // LeftPrecision -  округление левого валютного узла
	RightPrecision int64   `json:"right_precision"` // RightPrecision -  округление левого валютного узла
	TradeFee       float64 `json:"trade_fee"`       // TradeFee - комиссия при покупке/продаже
}

type Trade struct { // Trade - в будущем понадобится для автоматизации продажи, пока skip
	Executed    bool
	AmountLeft  Amount
	AmountRight Amount
	LeftFee     Amount
	RightFee    Amount
}

func NewBranch(
	market *Market,
	left *Node,
	right *Node,
	lrWeight float64,
	weightSize int64,
	exCount int64,
	exSymbol string,
	leftPrecision int64,
	rightPrecision int64,
	tradeFee float64,
	asks []*Lot,
	bids []*Lot,
) *Branch {
	return &Branch{
		Market:     market,
		Left:       left,
		Right:      right,
		LRWeight:   lrWeight,
		WeightSize: weightSize,
		ExCount:    exCount,
		ExSymbol:   exSymbol,
		BranchFilters: &BranchFilters{
			LeftPrecision:  leftPrecision,
			RightPrecision: rightPrecision,
			TradeFee:       tradeFee,
		},
		Asks: asks,
		Bids: bids,
	}
}

func (branch *Branch) ToString() string { // Строковое представление направления обмена
	return fmt.Sprintf("%s ---- > %s", branch.Left.Name, branch.Right.Name)
}
