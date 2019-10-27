package main

import "crypto/sha256"
import "fmt"
import "os"
import "time"

func main() {
	start := time.Now()
	if len(os.Args) == 2 {
		block_hash := os.Args[1]
		iter_hash := iterate([]byte(block_hash), uint64(8589934592)) // 2^33
		fmt.Printf("%x", iter_hash)
	} else {
		fmt.Println("Bad input")
	}
	elapsed := time.Since(start)
	fmt.Printf("\n Time elapsed: %s", elapsed)
}

func iterate(hash []byte, iter uint64) [32]byte {
	iter_hash := sha256.Sum256([]byte(hash))
	for i := uint64(1); i < iter; i++ {
		iter_hash = sha256.Sum256(iter_hash[:])
	}
	return iter_hash
}
