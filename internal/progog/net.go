package progog

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var maxRetries int

func SetMaxRetries(retries int) {
	if retries < 1 {
		maxRetries = 1
	} else {
		maxRetries = retries
	}
}

func tryGetBlock(hash string) (*Block, error) {
	var lastErr error

	url := "https://blockchain.info/rawblock/" + hash

	for retry := 0; retry < maxRetries; retry++ {
		resp, err := http.Get(url)
		if err != nil {
			lastErr = fmt.Errorf("HTTP request error on attempt %d: %v", retry+1, err)
			if retry < maxRetries-1 {
				time.Sleep(time.Duration(retry+1) * time.Second)
				continue
			}
			return nil, lastErr
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			lastErr = fmt.Errorf("HTTP status %d on attempt %d", resp.StatusCode, retry+1)
			if retry < maxRetries-1 {
				time.Sleep(time.Duration(retry+1) * time.Second)
				continue
			}
			return nil, lastErr
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			lastErr = fmt.Errorf("HTTP response read error on attempt %d: %v", retry+1, err)
			if retry < maxRetries-1 {
				time.Sleep(time.Duration(retry+1) * time.Second)
				continue
			}
			return nil, lastErr
		}

		var block Block
		if err := json.Unmarshal(body, &block); err != nil {
			lastErr = fmt.Errorf("JSON parsing error on attempt %d: %v", retry+1, err)
			if retry < maxRetries-1 {
				time.Sleep(time.Duration(retry+1) * time.Second)
				continue
			}
			return nil, lastErr
		}

		return &block, nil
	}

	return nil, fmt.Errorf("failed to fetch block %s after %d attempts: %v", hash, maxRetries, lastErr)
}
