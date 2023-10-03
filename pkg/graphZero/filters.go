package graphZero

func BranchFilterByLeft(left *string, branches []*Branch) (result []*Branch) {
	if left == nil {
		return branches
	}
	result = []*Branch{}
	for _, x := range branches {
		if x.Left.Name == *left {
			result = append(result, x)
		}
	}
	return
}

func BranchFilterByNode(node *string, branches []*Branch) (result []*Branch) {
	if node == nil {
		return branches
	}
	result = []*Branch{}
	for _, x := range branches {
		if x.Right.Name == *node || x.Left.Name == *node {
			result = append(result, x)
		}
	}
	return
}

func BranchFilterByRight(right *string, branches []*Branch) (result []*Branch) {
	if right == nil {
		return branches
	}
	result = []*Branch{}
	for _, x := range branches {
		if x.Right.Name == *right {
			result = append(result, x)
		}
	}
	return
}

func BranchFilterBySymbol(symbol *string, branches []*Branch) (result []*Branch) {
	if symbol == nil {
		return branches
	}
	result = []*Branch{}
	for _, x := range branches {
		if x.ExSymbol == *symbol {
			result = append(result, x)
		}
	}
	return
}

func BranchFilterByMarket(marketList []*Market, branches []*Branch) (answer KernelMarketPrice) {
	answer = KernelMarketPrice{[]KernelPriceList{}}
	var appendStatus bool
	for _, x := range branches {
		if len(marketList) != 0 {
			for _, k := range marketList {
				if x.Market == k {
					appendStatus = false
					for i, _ := range answer.Markets {
						if answer.Markets[i].Market == x.Market.Name {
							answer.Markets[i].Data = append(answer.Markets[i].Data, KernelPrice{
								Symbol:       x.ExSymbol,
								Left:         x.Left.Name,
								Right:        x.Right.Name,
								Price:        x.LRWeight,
								PriceQ:       x.LRWeightQ,
								AveragePrice: x.AveragePrice,
							})
							appendStatus = true
						}
					}
					if !appendStatus {
						answer.Markets = append(answer.Markets, KernelPriceList{
							Market: x.Market.Name,
							Data: []KernelPrice{{
								Symbol:       x.ExSymbol,
								Left:         x.Left.Name,
								Right:        x.Right.Name,
								Price:        x.LRWeight,
								PriceQ:       x.LRWeightQ,
								AveragePrice: x.AveragePrice,
							}},
						})
					}
				}
			}
		} else {
			appendStatus = false
			for i, _ := range answer.Markets {
				if answer.Markets[i].Market == x.Market.Name {
					answer.Markets[i].Data = append(answer.Markets[i].Data, KernelPrice{
						Symbol:       x.ExSymbol,
						Left:         x.Left.Name,
						Right:        x.Right.Name,
						Price:        x.LRWeight,
						PriceQ:       x.LRWeightQ,
						AveragePrice: x.AveragePrice,
					})
					appendStatus = true
				}
			}
			if !appendStatus {
				answer.Markets = append(answer.Markets, KernelPriceList{
					Market: x.Market.Name,
					Data: []KernelPrice{{
						Symbol:       x.ExSymbol,
						Left:         x.Left.Name,
						Right:        x.Right.Name,
						Price:        x.LRWeight,
						PriceQ:       x.LRWeightQ,
						AveragePrice: x.AveragePrice,
					}},
				})
			}
		}
	}
	return
}
