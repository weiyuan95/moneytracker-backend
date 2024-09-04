package coingate

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type CurrencyRateItem struct {
	ttl  time.Time
	rate string
}

type CurrencyRateOptions struct {
	From     string
	To       string
	CacheFor time.Duration
}

// We're not too concerned about cache hits/misses from concurrent function calls
// since it's fine to just make another call to the API. If we really care about
// cache consistency, we can use a mutex to lock the cache map.
var inMemoryCurrencyRateCache = make(map[[2]string]CurrencyRateItem)

func CurrencyRate(opts CurrencyRateOptions) (string, error) {
	cachedRateItem, present := inMemoryCurrencyRateCache[[2]string{opts.From, opts.To}]

	if present {
		now := time.Now()

		if now.Before(cachedRateItem.ttl) {
			return cachedRateItem.rate, nil
		}
	}

	// Item is not cached, or is expired, so we need to re-call the API

	url := fmt.Sprintf("https://api.coingate.com/v2/rates/merchant/%s/%s", opts.From, opts.To)
	resp, err := http.Get(url)
	// Arbitrarily wait for 0.5s to avoid spamming the public API, regardless
	// if there's an error or not.
	time.Sleep(500 * time.Millisecond)

	// We don't really care if there's an error in closing the response body
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if err != nil {
		return "-1", err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "-1", err
	}

	if len(body) == 0 {
		return "-1", fmt.Errorf("empty response")
	}

	stringedBody := string(body)

	inMemoryCurrencyRateCache[[2]string{opts.From, opts.To}] = CurrencyRateItem{
		time.Now().Add(opts.CacheFor),
		stringedBody,
	}

	return stringedBody, nil
}
