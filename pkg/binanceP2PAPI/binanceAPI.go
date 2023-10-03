package binanceP2PAPI

import (
	"MarketEye/config"
	"MarketEye/pkg/fhttp"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"sync"
	"time"
)

type Manager struct {
	cfg           *config.Config
	mu            *sync.RWMutex
	fhttpClient   *fhttp.Client
	domainHistory map[string]time.Time
}

func NewBinanceP2PAPI(cfg *config.Config, fhttpClient *fhttp.Client) *Manager {
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

func (b *Manager) OrderBook(asset string, fiat string, banks []string, tradeType string, rows int64) (*PriceResponse, error) {
	url := endpointP2PPrice
	b.RequestDelay("0.0.0.0", url, delayPrice)
	method := fhttp.MethodGET
	var responseBody, requestBody []byte
	var statusCode int
	var headers = make(map[string]string)
	headers["Accept"] = "*/*"
	headers["Accept-Encoding"] = "gzip, deflate, br"
	headers["Accept-Language"] = "ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3"
	headers["Cache-Control"] = "no-cache"
	headers["Connection"] = "keep-alive"
	headers["Content-Length"] = "123"
	headers["Host"] = "p2p.binance.com"
	headers["Origin"] = "https://p2p.binance.com"
	headers["Pragma"] = "no-cache"
	headers["TE"] = "Trailers"
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0"
	params := PriceBody{
		Asset:         asset,
		Fiat:          fiat,
		MerchantCheck: false,
		Page:          1,
		PayTypes:      banks,
		PublisherType: nil,
		Rows:          rows,
		TradeType:     tradeType,
	}
	requestBody, _ = json.Marshal(params)
	responseBody, statusCode, err := b.fhttpClient.Request(method, url, requestBody, nil, headers)
	if err != nil {
		return nil, err
	}
	var orderBook PriceResponse
	err = json.Unmarshal(responseBody, &orderBook)
	if err != nil {
		return nil, errors.Wrap(err, "Binance P2P API response error: cannot unmarshal")
	}
	if statusCode != 200 {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot reach the service: status code %d", statusCode))
	}
	return &orderBook, nil
}
