package main

import (
	"math/rand"
	"math/big"
	"fmt"
)

func main() {
	n, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("random number: %d\n", n.Int64())
	
	 v := "temp"
        for {
                v := "temp2"
                fmt.Println(v)
                break
        }
        fmt.Println(v)


}
