package main

import "fmt"
import "github.com/minio/sha256-simd"
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

func iterate(hash []byte, iter int) []byte {
	hasher := sha256.New()
	for i := 1; i < iter; i++ {
		hasher.Reset()
		hasher.Write(hash)
		hash = hasher.Sum([]byte{})
	}
	return hash
}
