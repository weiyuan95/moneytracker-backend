package coingate

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// NOTE: Go tests are run in parallel, and since it's a singleton cache, running all
// the tests at once isn't thread safe. This is why we have to run the tests with specific From/To currency symbols
// to avoid cache hits/misses from other tests.
// When this is actually running, it might result in some cache misses, but it shouldn't be a problem
// at such a small scale

func TestCurrencyRateCacheHit(t *testing.T) {
	_, err := CurrencyRate(CurrencyRateOptions{"BTC", "USD", 5 * time.Minute})
	assert.NoError(t, err)

	start := time.Now()
	_, err = CurrencyRate(CurrencyRateOptions{"BTC", "USD", 5 * time.Minute})
	assert.NoError(t, err)

	// In an event of a cache hit, should take < 10ms to return
	assert.Less(t, time.Since(start).Milliseconds(), int64(10))
}

func TestCurrencyRateCacheMiss(t *testing.T) {
	start := time.Now()
	_, err := CurrencyRate(CurrencyRateOptions{"BTC", "ETH", 5 * time.Minute})
	assert.NoError(t, err)

	_, err = CurrencyRate(CurrencyRateOptions{"ETH", "BTC", 5 * time.Minute})
	assert.NoError(t, err)

	// In an event of a cache miss, should take > 1000ms to return since that is the throttle time
	// 500ms + 500ms for each request
	assert.Greater(t, time.Since(start).Milliseconds(), int64(1000))
}

func TestCurrencyRateWrongSymbol(t *testing.T) {
	rate, err := CurrencyRate(CurrencyRateOptions{"BTC", "ABCD", 5 * time.Minute})
	assert.Error(t, err)
	assert.Equal(t, rate, "-1")

	rate, err = CurrencyRate(CurrencyRateOptions{"ABCD", "ETH", 5 * time.Minute})
	assert.Error(t, err)
	assert.Equal(t, rate, "-1")
}

func TestCurrencyRateCacheFor(t *testing.T) {
	start := time.Now()
	_, err := CurrencyRate(CurrencyRateOptions{"BTC", "SGD", 5 * time.Second})
	assert.NoError(t, err)

	time.Sleep(5 * time.Second)

	_, err = CurrencyRate(CurrencyRateOptions{"BTC", "SGD", 5 * time.Second})
	assert.NoError(t, err)

	// In an event of a cache miss, should take > 5500ms to return since that is the throttle time + cache time
	assert.Greater(t, time.Since(start).Milliseconds(), int64(5500))
}
