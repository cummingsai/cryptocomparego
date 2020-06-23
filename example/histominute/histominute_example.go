package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cummingsai/cryptocomparego"
)

func main() {
	client := cryptocomparego.NewClient("4a0867ab22e8806d04f9bf19a88c658d26d1fb7d4753a28536299a452f96f441", nil)
	ctx := context.TODO()

	histominRequest := cryptocomparego.NewHistominuteRequest("BTC", "USDT", "2014-01-01")
	data, _, err := client.Histomin.Get(ctx, histominRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
