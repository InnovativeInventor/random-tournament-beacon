package main

import "crypto/sha256"
import "fmt"
import "os"

func main() {
	if len(os.Args) == 2 {
		block_hash := os.Args[1]
		iter_hash := iterate([]byte(block_hash), 10000000)
		fmt.Printf("%x", iter_hash)
	} else {
		fmt.Println("Bad input")
	}
}

func iterate(hash []byte, iter int) [32]byte {
	iter_hash := sha256.Sum256([]byte(hash))
	for i := 1; i < iter; i++ {
		iter_hash = sha256.Sum256(iter_hash[:])
	}
	return iter_hash
}
