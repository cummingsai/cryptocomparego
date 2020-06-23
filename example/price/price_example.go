package main

import (
	"fmt"

	"github.com/cummingsai/cryptocomparego"
	"github.com/cummingsai/cryptocomparego/context"
)

func main() {

	client := cryptocomparego.NewClient("4a0867ab22e8806d04f9bf19a88c658d26d1fb7d4753a28536299a452f96f441", nil)
	ctx := context.TODO()

	priceRequest := cryptocomparego.NewPriceRequest("BTC", []string{"EUR"})
	priceList, _, err := client.Price.List(ctx, priceRequest)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n", err)
	}

	for _, coin := range priceList {
		fmt.Printf("Coin %s - %f\n", coin.Name, coin.Value)
	}
}
