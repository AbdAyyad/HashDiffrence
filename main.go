package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

type Hash [32]byte

func findDifference(firstHash Hash, secondHash Hash) int {
	counter := 0
	for i := 0; i < len(firstHash); i++ {
		for j := 0; j < 8; j++ {
			bit1 := uint(firstHash[i]) & (1 << uint(j))
			bit2 := uint(secondHash[i]) & (1 << uint(j))

			if bit1 != bit2 {
				counter++
			}
		}
	}
	return counter
}

func main() {
	c1 := sha256.Sum256([]byte(os.Args[1]))
	c2 := sha256.Sum256([]byte(os.Args[2]))
	diff := findDifference(c1, c2)
	fmt.Printf("first hash: %3v\nsecond hash:%3v\n", c1, c2)
	fmt.Printf("the diffrence in bits: %v\n", diff)
}
