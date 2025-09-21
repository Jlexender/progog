package main

import (
	"flag"
	"fmt"
	"jex/progog/internal/progog"
	"os"
)

func main() {
	startHash := flag.String("start", "", "Starting block hash")
	depth := flag.Int("depth", 1, "Depth of traversal")
	flag.Parse()
	
	if *startHash == "" || *depth < 1 {
		throwExit()
	}
	
	progog.Start(*startHash, *depth)
}


func throwExit() {
	fmt.Println("Usage: progog -start <block_hash> -depth <depth>")
	os.Exit(1)
}

