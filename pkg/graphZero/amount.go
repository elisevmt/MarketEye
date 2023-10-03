package graphZero

type Amount struct {
	Market *Market // Ссылка на биржу, на которой хранится объём средств
	Value  float64 // Количество валюты
	Node   *Node   // Валютный узел (ETH, BTC, ...)
}

func NewAmount(
	market *Market,
	value float64,
	node *Node,
) *Amount {
	return &Amount{
		Market: market,
		Value:  value,
		Node:   node,
	}
}

func (amount *Amount) Move(
	branch *Branch,
) {
	amount.Value = amount.Value * branch.LRWeight / branch.BranchFilters.TradeFee // Move - продажа, конвертация валюты по ветке Branch
	amount.Node = branch.Right                                                    // Учитывается комиссия продажи и актуальный price
}

func (amount *Amount) PredictMove(
	branch *Branch,
) *Amount { // PredictMove - это условная продажа, нужна для того, чтобы проверить исход манипуляции на бирже
	copyAmount := *amount
	copyAmount.Move(branch)
	return &copyAmount
}

func (amount *Amount) WayExecute( // WayExecute - исполняет список Branch над объёмом средств (цепочка продаж)
	way []*Branch,
) {
	for _, el := range way {
		amount.Move(el)
	}
}

func NodeContains(
	amountList []*Amount,
	node *Node,
) bool {
	for _, r := range amountList {
		if r.Node == node {
			return true
		}
	}
	return false
}
