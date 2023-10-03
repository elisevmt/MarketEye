package binanceAPI

import (
	"MarketEye/config"
	"MarketEye/pkg/fhttp"
	"MarketEye/pkg/graphZero"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Manager struct {
	cfg           *config.Config
	mu            *sync.RWMutex
	fhttpClient   *fhttp.Client
	domainHistory map[string]time.Time
}

func NewBinanceAPI(cfg *config.Config, fhttpClient *fhttp.Client) *Manager {
	mu := sync.RWMutex{}
	return &Manager{
		cfg:           cfg,
		mu:            &mu,
		fhttpClient:   fhttpClient,
		domainHistory: make(map[string]time.Time),
	}
}

func (b *Manager) RequestDelay(domenProxy string, url string, duration time.Duration) {
	var wait time.Duration
	b.mu.Lock()
	if lastUpdate, ok := b.domainHistory[domenProxy+url]; ok {
		currentTime := time.Now()
		lastUpdate = lastUpdate.Add(duration)

		if currentTime.Before(lastUpdate) { //wait
			wait = lastUpdate.Sub(currentTime)
			b.domainHistory[domenProxy+url] = lastUpdate.Add(duration)
		} else {
			b.domainHistory[domenProxy+url] = (time.Now()).Add(duration)
		}
	} else {
		b.domainHistory[domenProxy+url] = (time.Now()).Add(duration)
	}
	b.mu.Unlock()
	time.Sleep(wait)
}

func (b *Manager) Auth(outString string) string {
	h := hmac.New(sha256.New, []byte(b.cfg.BinanceAccount.APISecret))
	h.Write([]byte(outString))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func (b *Manager) GetTicker() ([]Ticker, error) {
	url := endpointTickerPrice
	b.RequestDelay("0.0.0.0", url, delayTickerPrice)
	method := fhttp.MethodGET
	var responseBody, requestBody []byte
	var statusCode int

	var headers = make(map[string]string)

	responseBody, statusCode, err := b.fhttpClient.Request(method, url, requestBody, nil, headers)
	if err != nil {
		return nil, err
	}

	var ticker []Ticker

	err = json.Unmarshal(responseBody, &ticker)
	if err != nil {
		return nil, errors.Wrap(err, "Binance API response error: cannot unmarshal")
	}
	if statusCode != 200 {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot reach the service: status code %d", statusCode))
	}
	return ticker, nil
}

func (b *Manager) GetTicker24hr() ([]Ticker24hr, error) {
	url := endpointTicker24hr
	b.RequestDelay("0.0.0.0", url, delayTicker24hr)
	method := fhttp.MethodGET
	var responseBody, requestBody []byte
	var statusCode int

	var headers = make(map[string]string)

	responseBody, statusCode, err := b.fhttpClient.Request(method, url, requestBody, nil, headers)
	if err != nil {
		return nil, err
	}

	var ticker []Ticker24hr

	err = json.Unmarshal(responseBody, &ticker)
	if err != nil {
		return nil, errors.Wrap(err, "Binance API response error: cannot unmarshal")
	}
	if statusCode != 200 {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot reach the service: status code %d", statusCode))
	}
	return ticker, nil
}

func (b *Manager) GetExchangeInfo() (*ExchangeInfo, error) {
	url := endpointExchangeInfo
	b.RequestDelay("0.0.0.0", url, delayExchangeInfo)
	method := fhttp.MethodGET
	var responseBody, requestBody []byte
	var statusCode int

	var headers = make(map[string]string)

	responseBody, statusCode, err := b.fhttpClient.Request(method, url, requestBody, nil, headers)
	if err != nil {
		return nil, err
	}

	var exchangeInfo ExchangeInfo

	err = json.Unmarshal(responseBody, &exchangeInfo)
	if err != nil {
		return nil, errors.Wrap(err, "Binance API response error: cannot unmarshal")
	}
	if statusCode != 200 {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot reach the service: status code %d", statusCode))
	}
	return &exchangeInfo, nil
}

func (b *Manager) OrderBook(symbol string, limit int64) (*SourceBook, error) {
	url := endpointOrderBook
	b.RequestDelay("0.0.0.0", url, delayOrderBook)
	method := fhttp.MethodGET
	var responseBody, requestBody []byte
	var statusCode int

	var headers = make(map[string]string)
	var query = make(map[string]string)
	query["symbol"] = symbol
	query["limit"] = strconv.Itoa(int(limit))
	responseBody, statusCode, err := b.fhttpClient.Request(method, url, requestBody, query, headers)
	if err != nil {
		return nil, err
	}

	var orderBook SourceBook

	err = json.Unmarshal(responseBody, &orderBook)
	if err != nil {
		return nil, errors.Wrap(err, "Binance API response error: cannot unmarshal")
	}
	if statusCode != 200 {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot reach the service: status code %d", statusCode))
	}
	return &orderBook, nil
}

func (b *Manager) OrderBookTransform(o *SourceBook) *graphZero.OrderBook {
	var orderBook graphZero.OrderBook
	for _, x := range o.Asks {
		xPrice, _ := strconv.ParseFloat(x[0], 10)
		xAmount, _ := strconv.ParseFloat(x[1], 10)
		orderBook.Asks = append(orderBook.Asks, &graphZero.Lot{
			Price:  xPrice,
			Amount: xAmount,
		})
	}
	for _, x := range o.Bids {
		xPrice, _ := strconv.ParseFloat(x[0], 10)
		xAmount, _ := strconv.ParseFloat(x[1], 10)
		orderBook.Bids = append(orderBook.Bids, &graphZero.Lot{
			Price:  xPrice,
			Amount: xAmount,
		})
	}
	return &orderBook
}

func (b *Manager) OrderBookTransformWSS(o *OrderBookWSS) *graphZero.OrderBook {
	var orderBook graphZero.OrderBook
	for _, x := range o.Asks {
		xPrice, _ := strconv.ParseFloat(x[0], 10)
		xAmount, _ := strconv.ParseFloat(x[1], 10)
		orderBook.Asks = append(orderBook.Asks, &graphZero.Lot{
			Price:  xPrice,
			Amount: xAmount,
			PriceQ: xPrice,
		})
	}
	for _, x := range o.Bids {
		xPrice, _ := strconv.ParseFloat(x[0], 10)
		xAmount, _ := strconv.ParseFloat(x[1], 10)
		orderBook.Bids = append(orderBook.Bids, &graphZero.Lot{
			Price:  xPrice,
			Amount: xAmount,
			PriceQ: xPrice,
		})
	}
	return &orderBook
}

func (b *Manager) GetServerTime() (*ServerTime, error) {
	url := endpointServerTime
	b.RequestDelay("0.0.0.0", url, delayServerTIme)
	method := fhttp.MethodGET
	var responseBody, requestBody []byte
	var statusCode int

	var headers = make(map[string]string)
	responseBody, statusCode, err := b.fhttpClient.Request(method, url, requestBody, nil, headers)
	if err != nil {
		return nil, err
	}

	var serverTime ServerTime

	err = json.Unmarshal(responseBody, &serverTime)
	if err != nil {
		return nil, errors.Wrap(err, "Binance API response error: cannot unmarshal")
	}
	if statusCode != 200 {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot reach the service: status code %d", statusCode))
	}
	return &serverTime, nil
}

func (b *Manager) GetTradeFee() ([]*TradeFee, error) {
	url := endpointTradeFee
	b.RequestDelay("0.0.0.0", url, delayTradeFee)
	method := fhttp.MethodGET
	var responseBody, requestBody []byte
	var statusCode int
	var payload string
	serverTime, _ := b.GetServerTime()
	payload = fmt.Sprintf("timestamp=%v&recvWindow=%d", serverTime.Time, 5000)
	requestBody = []byte(fmt.Sprintf("%s&signature=%s", payload, b.Auth(payload)))
	var headers = make(map[string]string)
	headers["X-MBX-APIKEY"] = b.cfg.BinanceAccount.APIKey
	url = fmt.Sprintf("%s?%s", url, requestBody)
	responseBody, statusCode, err := b.fhttpClient.Request(method, url, nil, nil, headers)
	if err != nil {
		return nil, err
	}
	//fmt.Println("Trade fee response body: ", string(responseBody))
	var tradeFee []*TradeFee
	err = json.Unmarshal(responseBody, &tradeFee)
	if statusCode != 200 {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot reach the service: status code %d", statusCode))
	}
	if err != nil {
		return nil, errors.Wrap(err, "Binance API response error: cannot unmarshal")
	}

	return tradeFee, nil
}

func (b *Manager) bncAPIRequest(method, endpoint, payload string, sign bool) (response []byte, err error) { // TODO: добавить обработку ошибок в json ответа от binance
	client := &http.Client{}
	url := endpoint
	if sign {
		time, err := b.GetServerTime()
		if err != nil {
			return nil, err
		}
		payload = fmt.Sprintf("%s&timestamp=%v&recvWindow=%d", payload, time.Time, 1000) // TODO: разобраться с этим
		mac := hmac.New(sha256.New, []byte(b.cfg.BinanceAccount.APISecret))
		_, err = mac.Write([]byte(payload))
		if err != nil {
			return nil, err
		}
		payload = fmt.Sprintf("%s&signature=%s", payload, hex.EncodeToString(mac.Sum(nil)))
	}

	var req *http.Request
	switch method {
	case http.MethodGet: // TODO: проверить, можно ли сделать чище
		req, err = http.NewRequest(method, fmt.Sprintf("%s?%s", url, payload), nil)
		if err != nil {
			return nil, err
		}
	case http.MethodPost:
		req, err = http.NewRequest(method, fmt.Sprintf("%s", url), strings.NewReader(payload))
		if err != nil {
			return nil, err
		}
		req.Header.Add("Content-QuantityType", "application/x-www-form-urlencoded")
	default:
		return nil, errors.New("unknown method")
	}
	if sign {
		req.Header.Add("X-MBX-APIKEY", b.cfg.BinanceAccount.APIKey)
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(response))
	}
	return response, err
}

func (b *Manager) MakeLimitOrder(symbol, side, quantity, price string) (NewOrderResponseShort, error) {
	var orderResponse NewOrderResponseShort
	if symbol == "" || side == "" || quantity == "" {
		return orderResponse, errors.New("not enough arguments")
	}
	params := make(map[string]string)
	params["symbol"] = symbol
	params["side"] = side
	params["quantity"] = quantity
	params["type"] = "LIMIT"
	params["price"] = price
	params["newOrderRespType"] = "FULL"
	params["timeInForce"] = "IOC"
	var payload string
	var i int
	for k, v := range params {
		if i == 0 {
			payload += k + "=" + v
			i++
			continue
		}
		payload += "&" + k + "=" + v
	}
	var orderResponseFull NewOrderResponse
	response, err := b.bncAPIRequest("POST", endpointOrder, payload, true)
	if err != nil {
		return orderResponse, err
	}
	err = json.Unmarshal(response, &orderResponseFull) // TODO: fatal
	if err != nil {
		return orderResponse, err
	}

	var prices, count float64
	for _, r := range orderResponseFull.Fills {
		price, err := strconv.ParseFloat(r.Price, 64)
		if err != nil {
			return orderResponse, err
		}
		prices += price
		count++
	}
	origQty, err := strconv.ParseFloat(orderResponseFull.OrigQty, 64)
	if err != nil {
		return orderResponse, err
	}
	cummulativeQuoteQty, err := strconv.ParseFloat(orderResponseFull.CummulativeQuoteQty, 10)
	if err != nil {
		return orderResponse, err
	}
	orderResponse.AVGPrice = prices / count
	orderResponse.OrigQty = origQty
	orderResponse.CummulativeQuoteQty = cummulativeQuoteQty
	return orderResponse, err
}

func (b *Manager) GetWSSConn(symbol string) (*websocket.Conn, error) {
	symbol = strings.ToLower(symbol)
	//url := fmt.Sprintf("wss://stream.binance.com/ws/%s@price/%s@depth", symbol, symbol)
	url := fmt.Sprintf("wss://stream.binance.com/ws/%s@depth/%s@ticker", symbol, symbol)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return conn, err
}

func (b *Manager) GetBatchWSSCOnn(symbols []string) ([]*websocket.Conn, error) {
	batchUrlList := []string{""}
	k := 0
	for i, symbol := range symbols {
		if i > 199*k+1 {
			k = k + 1
			batchUrlList = append(batchUrlList, "")
		}
		symbol = strings.ToLower(symbol)
		batchUrlList[k] += fmt.Sprintf("%s@depth/", symbol)
	}
	var connections []*websocket.Conn
	for _, batchUrl := range batchUrlList {
		batchUrl = batchUrl[:len(batchUrl)-1]
		url := fmt.Sprintf("wss://stream.binance.com/stream?streams=%s", batchUrl)
		conn, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			return nil, err
		}
		connections = append(connections, conn)
	}
	return connections, nil
}
