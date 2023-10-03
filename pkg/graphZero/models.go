package graphZero

type MarginResult struct { // MarginResult пока не трогаем, понадобится для автоматизации позже
	Way []*Branch
	AmountHistory []*Amount
	Margin float64
}
