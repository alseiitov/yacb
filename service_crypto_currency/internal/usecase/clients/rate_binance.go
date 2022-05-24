package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alseiitov/yacb/service_crypto_currency/internal/entity"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type binanceClient struct {
	URL        string
	HTTPClient *http.Client
}

func NewBinanceClient(url string) *binanceClient {

	transport := &http.Transport{
		TLSHandshakeTimeout: 30 * time.Second,
	}

	httpClient := &http.Client{
		Transport: transport,
		Timeout:   1 * time.Minute,
	}

	return &binanceClient{
		URL:        url,
		HTTPClient: httpClient,
	}
}

type binanceRateResp struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type binanceRateErrorResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

func (c *binanceClient) GetCurrentRate(ctx context.Context, symbol string) (entity.Rate, error) {

	var values = url.Values{}

	symbolPair := fmt.Sprintf("%sUSDT", symbol)
	values.Set("symbol", symbolPair)

	URL := fmt.Sprintf("%s/ticker/price?%s", c.URL, values.Encode())

	request, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return entity.Rate{}, err
	}

	resp, err := c.HTTPClient.Do(request)
	if err != nil {
		return entity.Rate{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp binanceRateErrorResp
		err = json.NewDecoder(resp.Body).Decode(&errResp)
		if err != nil {
			return entity.Rate{}, err
		}

		return entity.Rate{}, errors.New(errResp.Msg)
	}

	var rateResp binanceRateResp

	err = json.NewDecoder(resp.Body).Decode(&rateResp)
	if err != nil {
		return entity.Rate{}, err
	}

	now := time.Now()

	price, err := strconv.ParseFloat(rateResp.Price, 32)
	if err != nil {
		return entity.Rate{}, err
	}

	rate := entity.Rate{
		CryptoCurrencyName: symbol,
		Price:              float32(price),
		Date:               &now,
	}

	return rate, nil
}
