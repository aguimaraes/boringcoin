package main

import (
	"fmt"
	"github.com/aguimaraes/boringcoin/pkg/block"
)

func main() {
	c := []block.Block{block.Genesis()}
	pb := c[0]
	l := 20
	for i := 0; i < l; i++ {
		a := block.Next(pb)
		c = append(c, a)
		pb = a
		fmt.Printf("Block #%d\n", a.Index)
		fmt.Printf("Hash: %x\n\n", a.Hash.Sum(nil))
	}

}
