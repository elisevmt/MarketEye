package usecase

import (
	"MarketEye/config"
	"MarketEye/internal/kernel"
	"MarketEye/internal/models"
	"MarketEye/pkg/graphZero"
)

type kernelUC struct {
	cfg        *config.Config
	kernelTree *graphZero.Tree
	marketMap  *map[string]*graphZero.Market
}

func NewKernelUseCase(cfg *config.Config, kernelTree *graphZero.Tree, marketMap *map[string]*graphZero.Market) kernel.UseCase {
	return &kernelUC{
		cfg:        cfg,
		kernelTree: kernelTree,
		marketMap:  marketMap,
	}
}

func (k *kernelUC) FetchPrices(params models.FetchPricesParams) graphZero.KernelMarketPrice {
	var marketList []*graphZero.Market
	for _, v := range *k.marketMap {
		for _, r := range params.MarketList {
			if *r == v.Name {
				marketList = append(marketList, v)
			}
		}
	}
	return k.kernelTree.GetTreeKernel(marketList, params.Node, params.Left, params.Right, params.Symbol)
}

func (k *kernelUC) FetchPricesAverage(params models.FetchPricesParams) graphZero.KernelMarketPrice {
	var marketList []*graphZero.Market
	for _, v := range *k.marketMap {
		for _, r := range params.MarketList {
			if *r == v.Name {
				marketList = append(marketList, v)
			}
		}
	}
	return k.kernelTree.GetTreeKernel(marketList, params.Node, params.Left, params.Right, params.Symbol)
}

func (k *kernelUC) FetchOrderBook(params models.FetchOrderBookParams) graphZero.OrderBook {
	branch, err := k.kernelTree.GetBranch(params.Symbol, (*k.marketMap)[params.Market])
	if err != nil {
		return graphZero.OrderBook{Bids: nil, Asks: nil}
	}
	if params.Depth != nil {
		indexBids, indexAsks := 0, 0
		if len(branch.Bids)-1 >= int(*params.Depth) {
			indexBids = int(*params.Depth)
		} else if len(branch.Bids) != 0 {
			indexBids = len(branch.Bids) - 1
		}
		if len(branch.Asks)-1 >= int(*params.Depth) {
			indexAsks = int(*params.Depth)
		} else if len(branch.Asks) != 0 {
			indexAsks = len(branch.Asks) - 1
		}
		return graphZero.OrderBook{
			Bids: branch.Bids[0:indexBids],
			Asks: branch.Asks[0:indexAsks],
		}
	}
	return graphZero.OrderBook{
		Bids: branch.Bids,
		Asks: branch.Asks,
	}
}

func (k *kernelUC) FetchMarketList() (result models.FetchMarketListResponse) {
	result = models.FetchMarketListResponse{MarketList: []string{}}
	for m, _ := range *k.marketMap {
		result.MarketList = append(result.MarketList, m)
	}
	return
}
