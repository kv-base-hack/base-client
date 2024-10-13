package kaivestbinance

import (
	"net/http"
	"strconv"
	"time"

	binance "github.com/adshao/go-binance/v2"
	bf "github.com/adshao/go-binance/v2/futures"
	"github.com/kv-base-hack/common/httpclient"
)

const defaultClientTimeout = time.Second * 30

type KaivestBinanceClient struct {
	addr   string
	client *http.Client
}

func NewKaivestBinanceClient(addr string) *KaivestBinanceClient {
	httpClient := &http.Client{
		Timeout: defaultClientTimeout,
		Transport: &http.Transport{
			IdleConnTimeout:       time.Second * 120,
			ResponseHeaderTimeout: time.Second * 10,
		},
	}
	return &KaivestBinanceClient{
		addr:   addr,
		client: httpClient,
	}
}

func (obc *KaivestBinanceClient) GetFutureKLine(pair string, interval string, limit int, ts int64) ([]bf.Kline, error) {
	var klines []bf.Kline
	req, err := httpclient.NewRequest(http.MethodGet, obc.addr, "/binance/future/kline", nil, nil)
	if err != nil {
		return []bf.Kline{}, err
	}
	q := req.URL.Query()
	q.Add("pair", pair)
	q.Add("interval", interval)
	q.Add("limit", strconv.Itoa(limit))
	if ts > 0 {
		q.Add("end_ts", strconv.Itoa(int(ts)))
	}
	req.URL.RawQuery = q.Encode()
	_, err = httpclient.DoHTTPRequest(obc.client, req, &klines)
	if err != nil {
		return []bf.Kline{}, err
	}
	return klines, nil
}

func (obc *KaivestBinanceClient) GetFundingRate() (Premium, error) {
	var premium Premium
	req, err := httpclient.NewRequest(http.MethodGet, obc.addr, "/binance/future/premium", nil, nil)
	if err != nil {
		return Premium{}, err
	}
	_, err = httpclient.DoHTTPRequest(obc.client, req, &premium)
	if err != nil {
		return Premium{}, err
	}

	return premium, nil
}

func (obc *KaivestBinanceClient) GetOIStats(pair string, period string, limit int) ([]bf.OpenInterestStatistic, error) {
	var oiStats []bf.OpenInterestStatistic
	req, err := httpclient.NewRequest(http.MethodGet, obc.addr, "/binance/future/oi-stats", nil, nil)
	if err != nil {
		return []bf.OpenInterestStatistic{}, err
	}
	q := req.URL.Query()
	q.Add("pair", pair)
	q.Add("period", period)
	q.Add("limit", strconv.Itoa(limit))
	req.URL.RawQuery = q.Encode()
	_, err = httpclient.DoHTTPRequest(obc.client, req, &oiStats)
	if err != nil {
		return []bf.OpenInterestStatistic{}, err
	}

	return oiStats, nil
}

func (obc *KaivestBinanceClient) GetExchangeInfo() (bf.ExchangeInfo, error) {
	var exchangeInfo bf.ExchangeInfo
	req, err := httpclient.NewRequest(http.MethodGet, obc.addr, "/binance/future/exchange-info", nil, nil)
	if err != nil {
		return bf.ExchangeInfo{}, err
	}
	_, err = httpclient.DoHTTPRequest(obc.client, req, &exchangeInfo)
	if err != nil {
		return bf.ExchangeInfo{}, err
	}

	return exchangeInfo, nil
}

func (obc *KaivestBinanceClient) GetBookTicker(symbol string) ([]bf.BookTicker, error) {
	var bookTicker []bf.BookTicker
	req, err := httpclient.NewRequest(http.MethodGet, obc.addr, "/binance/future/oi-stats", nil, nil)
	if err != nil {
		return []bf.BookTicker{}, err
	}
	q := req.URL.Query()
	q.Add("symbol", symbol)
	req.URL.RawQuery = q.Encode()
	_, err = httpclient.DoHTTPRequest(obc.client, req, &bookTicker)
	if err != nil {
		return []bf.BookTicker{}, err
	}

	return bookTicker, nil
}

func (obc *KaivestBinanceClient) GetSpotKLine(pair string, interval string, limit int, ts int64) ([]binance.Kline, error) {
	var klines []binance.Kline
	req, err := httpclient.NewRequest(http.MethodGet, obc.addr, "/binance/spot/kline", nil, nil)
	if err != nil {
		return []binance.Kline{}, err
	}
	q := req.URL.Query()
	q.Add("pair", pair)
	q.Add("interval", interval)
	q.Add("limit", strconv.Itoa(limit))
	if ts > 0 {
		q.Add("start_ts", strconv.Itoa(int(ts)))
	}
	req.URL.RawQuery = q.Encode()
	_, err = httpclient.DoHTTPRequest(obc.client, req, &klines)
	if err != nil {
		return []binance.Kline{}, err
	}
	return klines, nil
}

func (obc *KaivestBinanceClient) GetAllCoinInfo() ([]CoinInfo, error) {
	var res []CoinInfo
	req, err := httpclient.NewRequest(http.MethodGet, obc.addr, "/binance/spot/all-coin-info", nil, nil)
	if err != nil {
		return []CoinInfo{}, err
	}
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	_, err = httpclient.DoHTTPRequest(obc.client, req, &res)
	if err != nil {
		return []CoinInfo{}, err
	}
	return res, nil
}

func (obc *KaivestBinanceClient) GetSpotBookTicker(symbols string) ([]binance.SymbolTicker, error) {
	var res []binance.SymbolTicker
	req, err := httpclient.NewRequest(http.MethodGet, obc.addr, "/binance/spot/book-ticker", nil, nil)
	if err != nil {
		return []binance.SymbolTicker{}, err
	}
	q := req.URL.Query()
	q.Add("symbols", symbols)
	req.URL.RawQuery = q.Encode()
	_, err = httpclient.DoHTTPRequest(obc.client, req, &res)
	if err != nil {
		return []binance.SymbolTicker{}, err
	}
	return res, nil
}

func (obc *KaivestBinanceClient) GetPairsWithUsdt() ([]string, error) {
	var res []string
	req, err := httpclient.NewRequest(http.MethodGet, obc.addr, "/binance/spot/spot-pair-with-usdt", nil, nil)
	if err != nil {
		return []string{}, err
	}
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	_, err = httpclient.DoHTTPRequest(obc.client, req, &res)
	if err != nil {
		return []string{}, err
	}
	return res, nil
}
