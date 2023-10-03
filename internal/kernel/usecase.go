package kernel

import (
	"MarketEye/internal/models"
	"MarketEye/pkg/graphZero"
)

type UseCase interface {
	FetchPrices(params models.FetchPricesParams) graphZero.KernelMarketPrice
	FetchPricesAverage(params models.FetchPricesParams) graphZero.KernelMarketPrice
	FetchOrderBook(params models.FetchOrderBookParams) graphZero.OrderBook
	FetchMarketList() models.FetchMarketListResponse
}
