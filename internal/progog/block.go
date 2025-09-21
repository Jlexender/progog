package progog

type Block struct {
	Hash      string   `json:"hash"`
	NextBlock []string `json:"next_block"`
}

var blockPool []*Block



