package progog

import (
	"fmt"
)

func Start(startHash string, depth int) {
	blockPool = make([]*Block, 0, depth)
	traverseInOrder(startHash, depth)
}

func traverseInOrder(hash string, depth int) {
	if depth == 0 {
		return
	}

	block, err := getBlock(hash)
	if err != nil {
		fmt.Printf("Error fetching block %s: %v\n", hash, err)
		return
	}

	blockPool = append(blockPool, block)

	for _, nextHash := range block.NextBlock {
		traverseInOrder(nextHash, depth-1)
	}

	fmt.Println(block.Hash)
}


