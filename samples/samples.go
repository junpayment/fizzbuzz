package main

import (
	"fmt"
	
	"github.com/junpayment/fizzbuzz"
)

func main() {
	sets := fizzbuzz.DoSync(100)
	for _, set := range sets {
		fmt.Printf("%d : %s\n", set.Input, set.Res)
	}
}
