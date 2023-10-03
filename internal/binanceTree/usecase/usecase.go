package usecase

import (
	"MarketEye/config"
	"MarketEye/internal/binanceTree"
	"MarketEye/pkg/binanceAPI"
	"MarketEye/pkg/fhttp"
	"MarketEye/pkg/graphZero"
	"MarketEye/pkg/logger"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/valyala/fastjson"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

type binanceTreeUC struct {
	cfg            *config.Config
	binanceManager *binanceAPI.Manager
	currencyTree   *graphZero.Tree
	market         *graphZero.Market
	logger         *logger.ApiLogger
}

func NewBinanceTreeUC(cfg *config.Config, market *graphZero.Market, kernelTree *graphZero.Tree, fhttp *fhttp.Client, logger *logger.ApiLogger) binanceTree.UseCase {
	return &binanceTreeUC{
		cfg:            cfg,
		binanceManager: binanceAPI.NewBinanceAPI(cfg, fhttp),
		currencyTree:   kernelTree,
		market:         market,
		logger:         logger,
	}
}

func (b *binanceTreeUC) BranchUpdate(branchRaw, branchReverse *graphZero.Branch) {
	conn, err := b.binanceManager.GetWSSConn(branchRaw.ExSymbol)
	if err != nil {
		log.Print("WSS conn creating error " + err.Error() + fmt.Sprintf(" %s", branchReverse.ExSymbol))
		b.logger.ErrorFull(err)
		time.Sleep(5)
		go b.BranchUpdate(branchRaw, branchReverse)
		return
	}
	for {
		_, message, readErr := conn.ReadMessage()
		if readErr != nil {
			b.logger.ErrorFull(readErr)
			_ = conn.Close()
			time.Sleep(5)
			go b.BranchUpdate(branchRaw, branchReverse)
			break
		}
		var p fastjson.Parser
		v, err := p.Parse(string(message))
		if err != nil {
			b.logger.ErrorFull(err)
		}
		socType := v.Get("e")
		notifyType := strings.ReplaceAll(socType.String(), "\"", "")
		if notifyType == "depthUpdate" {
			var orderBook binanceAPI.OrderBookWSS
			err := json.Unmarshal(message, &orderBook)
			if err != nil {
				b.logger.ErrorFull(err)
			}
			sourceBook := b.binanceManager.OrderBookTransformWSS(&orderBook)
			branchRaw.Bids = []*graphZero.Lot{}
			branchRaw.Asks = []*graphZero.Lot{}
			for i, _ := range sourceBook.Asks {
				if sourceBook.Asks[i].Amount != 0 {
					branchRaw.Asks = append(branchRaw.Asks, sourceBook.Asks[i])
				}
			}
			for i, _ := range sourceBook.Bids {
				if sourceBook.Bids[i].Amount != 0 {
					branchRaw.Bids = append(branchRaw.Bids, sourceBook.Bids[i])
				}
			}
			branchReverse.Asks = branchRaw.Asks
			branchReverse.Bids = branchRaw.Bids
			if len(sourceBook.Bids) != 0 {
				branchRaw.LRWeight = sourceBook.Bids[0].Price
				branchRaw.LeftWidth = sourceBook.Bids[0].Amount
				branchRaw.RightWidth = sourceBook.Bids[0].Price * sourceBook.Bids[0].Amount
				branchRaw.LRWeightQ = sourceBook.Bids[0].Price
			}
			if len(sourceBook.Asks) != 0 {
				branchReverse.LRWeight = 1 / sourceBook.Asks[0].Price
				branchReverse.LeftWidth = sourceBook.Asks[0].Price * sourceBook.Asks[0].Amount
				branchReverse.RightWidth = sourceBook.Asks[0].Amount
				branchReverse.LRWeightQ = sourceBook.Asks[0].Price
			}
		} else if notifyType == "24hrTicker" {
			var ticker binanceAPI.Ticker24hWSS
			err := json.Unmarshal(message, &ticker)
			if err != nil {
				b.logger.ErrorFull(err)
			}
			branchRaw.ExCount = int64(ticker.N)
			branchRaw.AveragePrice, _ = strconv.ParseFloat(ticker.WeightedAveragePrice, 64)
			branchReverse.ExCount = int64(ticker.N)
			branchReverse.AveragePrice = branchRaw.AveragePrice
		}
		conn.PongHandler()
	}
}

func (b *binanceTreeUC) BatchBranchUpdate(branchRawList, branchReverseList []*graphZero.Branch) {
	var exSymbols []string
	for _, el := range branchRawList {
		exSymbols = append(exSymbols, el.ExSymbol)
	}

	connections, err := b.binanceManager.GetBatchWSSCOnn(exSymbols)
	fmt.Println("Number of connections: ", len(connections), " - number of symbols: ", len(exSymbols))
	if err != nil {
		b.logger.ErrorFull(err)
	}
	for _, conn := range connections {
		go b.BatchConnectionResolver(conn)
	}
	//for {
	//	_, message, err := conn.ReadMessage()
	//	if err != nil {
	//		log.Fatalf(err.Error())
	//	}
	//	fmt.Println("____________________________________________")
	//	//fmt.Println(string(message))
	//	var orderBook binanceAPI.BatchOrderBookWSS
	//	err = json.Unmarshal(message, &orderBook)
	//	if err != nil {
	//		log.Fatalf(fmt.Sprintf("Cannot parse wss binance message to order book %s", err.Error()))
	//	}
	//	//fmt.Println(orderBook)
	//	fmt.Println(orderBook.Stream)
	//}
	//return
}

func (b *binanceTreeUC) BatchConnectionResolver(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			b.logger.ErrorFull(err)
		}
		//fmt.Println("____________________________________________")
		//fmt.Println(string(message))
		var orderBook binanceAPI.BatchOrderBookWSS
		err = json.Unmarshal(message, &orderBook)
		if err != nil {
			b.logger.ErrorFull(err)
		}
		//fmt.Println(orderBook)
		//fmt.Println(orderBook.Stream)
	}
	return
}

func (b *binanceTreeUC) TreeLoad() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	exchangeInfo, err := b.binanceManager.GetExchangeInfo()
	if err != nil {
		b.logger.ErrorFull(err)
	}
	tickers, err := b.binanceManager.GetTicker()
	if err != nil {
		b.logger.ErrorFull(err)
	}
	tickers24hr, err := b.binanceManager.GetTicker24hr()
	if err != nil {
		b.logger.ErrorFull(err)
	}
	tradeFee, err := b.binanceManager.GetTradeFee()
	if err != nil {
		b.logger.ErrorFull(err)
	}
	for _, symbol := range exchangeInfo.Symbols {
		if symbol.Status != "TRADING" {
			continue
		}
		nodeA := graphZero.NewNode(b.currencyTree.Root.Market, symbol.BaseAsset)
		nodeB := graphZero.NewNode(b.currencyTree.Root.Market, symbol.QuoteAsset)
		var weight float64
		for _, x := range tickers {
			if x.Symbol == symbol.Symbol {
				weight, _ = strconv.ParseFloat(x.Price, 10)
				break
			}
		}
		b.currencyTree.Append(nodeA)
		b.currencyTree.Append(nodeB)
		var leftPrecision int64
		var weightSize int64
		for _, y := range symbol.Filters {
			if y.FilterType == "LOT_SIZE" {
				step, _ := strconv.ParseFloat(y.StepSize, 10)
				leftPrecision = int64(math.Log10(step))
			}
			if y.FilterType == "PRICE_FILTER" {
				step, _ := strconv.ParseFloat(y.TickSize, 10)
				weightSize = int64(math.Log10(step))
			}
		}
		var exCount int64
		for _, r := range tickers24hr {
			if r.Symbol == symbol.Symbol {
				exCount = int64(r.Count)
				break
			}
		}
		var makerCommission float64
		var takerCommission float64
		for _, r := range tradeFee {
			if r.Symbol == symbol.Symbol {
				makerCommission, _ = strconv.ParseFloat(r.MakerCommission, 10)
				makerCommission += 1
				takerCommission, _ = strconv.ParseFloat(r.TakerCommission, 10)
				takerCommission += 1
				break
			}
		}
		branchRaw := b.currencyTree.Wire(nodeA, nodeB, weight*0, weightSize, exCount, symbol.Symbol, leftPrecision, leftPrecision, makerCommission, nil, nil)
		branchReverse := b.currencyTree.Wire(nodeB, nodeA, weight*0, weightSize, exCount, symbol.Symbol, leftPrecision, leftPrecision, takerCommission, nil, nil)
		go b.BranchUpdate(branchRaw, branchReverse)
	}
}

func (b *binanceTreeUC) BatchTreeLoad() {
	var branchRawList, branchReverseList []*graphZero.Branch
	exchangeInfo, err := b.binanceManager.GetExchangeInfo()
	if err != nil {
		b.logger.ErrorFull(err)
	}
	tickers, err := b.binanceManager.GetTicker()
	if err != nil {
		b.logger.ErrorFull(err)
	}
	tickers24hr, err := b.binanceManager.GetTicker24hr()
	if err != nil {
		b.logger.ErrorFull(err)
	}
	tradeFee, err := b.binanceManager.GetTradeFee()
	if err != nil {
		b.logger.ErrorFull(err)
	}
	for _, symbol := range exchangeInfo.Symbols {
		if symbol.Status != "TRADING" {
			continue
		}
		nodeA := graphZero.NewNode(b.currencyTree.Root.Market, symbol.BaseAsset)
		nodeB := graphZero.NewNode(b.currencyTree.Root.Market, symbol.QuoteAsset)
		var weight float64
		for _, x := range tickers {
			if x.Symbol == symbol.Symbol {
				weight, _ = strconv.ParseFloat(x.Price, 10)
				break
			}
		}
		b.currencyTree.Append(nodeA)
		b.currencyTree.Append(nodeB)
		var leftPrecision int64
		var weightSize int64
		for _, y := range symbol.Filters {
			if y.FilterType == "LOT_SIZE" {
				step, _ := strconv.ParseFloat(y.StepSize, 10)
				leftPrecision = int64(math.Log10(step))
			}
			if y.FilterType == "PRICE_FILTER" {
				step, _ := strconv.ParseFloat(y.TickSize, 10)
				weightSize = int64(math.Log10(step))
			}
		}
		var exCount int64
		for _, r := range tickers24hr {
			if r.Symbol == symbol.Symbol {
				exCount = int64(r.Count)
				break
			}
		}
		var makerCommission float64
		var takerCommission float64
		for _, r := range tradeFee {
			if r.Symbol == symbol.Symbol {
				makerCommission, _ = strconv.ParseFloat(r.MakerCommission, 10)
				makerCommission += 1
				takerCommission, _ = strconv.ParseFloat(r.TakerCommission, 10)
				takerCommission += 1
				break
			}
		}
		branchRaw := b.currencyTree.Wire(nodeA, nodeB, weight*0, weightSize, exCount, symbol.Symbol, leftPrecision, leftPrecision, makerCommission, nil, nil)
		branchReverse := b.currencyTree.Wire(nodeB, nodeA, weight*0, weightSize, exCount, symbol.Symbol, leftPrecision, leftPrecision, takerCommission, nil, nil)
		branchRawList = append(branchRawList, branchRaw)
		branchReverseList = append(branchReverseList, branchReverse)
		//go b.BranchUpdate(branchRaw, branchReverse)
	}
	go b.BatchBranchUpdate(branchRawList, branchReverseList)
}
