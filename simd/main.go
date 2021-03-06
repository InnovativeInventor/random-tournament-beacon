package main

import "fmt"
import "github.com/minio/sha256-simd"
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

func iterate(hash []byte, iter uint64) []byte {
	hasher := sha256.New()
	for i := uint64(1); i < iter; i++ {
		hasher.Reset()
		hasher.Write(hash)
		hash = hasher.Sum([]byte{})
	}
	return hash
}
