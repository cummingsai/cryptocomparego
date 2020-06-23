package main

import (
	"fmt"

	"github.com/cummingsai/cryptocomparego"
	"github.com/cummingsai/cryptocomparego/context"
)

func main() {

	client := cryptocomparego.NewClient("4a0867ab22e8806d04f9bf19a88c658d26d1fb7d4753a28536299a452f96f441", nil)
	ctx := context.TODO()

	socialStats, _, err := client.SocialStats.Get(ctx, 1182)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n", err)
	}

	fmt.Printf("Stats %+v\n", socialStats)
	fmt.Printf("General Name %s\n", socialStats.General.Name)

	for _, similarItem := range socialStats.CryptoCompare.SimilarItems {
		fmt.Printf("Similar Item %s\n", similarItem.Name)
	}
}
