package progog

type Block struct {
	Hash       string   `json:"hash"`
	NextBlock  []string `json:"next_block"`
	Version    int      `json:"ver"`
	MerkleRoot string   `json:"mrkl_root"`
	Nonce      uint32   `json:"nonce"`
}

var blockPool []*Block
