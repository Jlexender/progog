package progog

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getBlock(hash string) (*Block, error) {
	url := "https://blockchain.info/rawblock/" + hash

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("HTTP response read error: %v", err)
	}

	var block Block
	if err := json.Unmarshal(body, &block); err != nil {
		return nil, fmt.Errorf("JSON parsing error: %v", err)
	}

	

	return &block, nil
}