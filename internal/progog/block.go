package progog

type Block struct {
	Hash       string   `json:"hash"`
	NextBlock  []string `json:"next_block"`
	Version    int      `json:"ver"`
	MerkleRoot string   `json:"mrkl_root"`
	Nonce      uint32   `json:"nonce"`
	Time       uint32   `json:"time"`
	NTx        int      `json:"n_tx"`
	Size       int      `json:"size"`
	Height     int      `json:"height"`
	Weight     int      `json:"weight"`
	PrevBlock  string   `json:"prev_block"`
}

var BlockPool []*Block
