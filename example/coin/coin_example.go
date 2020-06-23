package main

import (
	"fmt"

	"github.com/cummingsai/cryptocomparego"
	"github.com/cummingsai/cryptocomparego/context"
)

func main() {

	client := cryptocomparego.NewClient("4a0867ab22e8806d04f9bf19a88c658d26d1fb7d4753a28536299a452f96f441", nil)
	ctx := context.TODO()

	coinList, _, err := client.Coin.List(ctx)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n", err)
	}

	for _, coin := range coinList {
		fmt.Printf("Coin %s - %s\n", coin.Name, coin.FullName)
	}
	fmt.Printf("A total of %d coins available\n", len(coinList))
}
