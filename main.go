package main

import (
	"fmt"
	"sync"

	"github.com/rjhoppe/go-cryptomasters/api"
)

// main goroutine
func main() {
	currencies := []string{"BTC", "ETH", "BHC"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		// lambda goroutine func
		// need to pass currency from outer func as local param
		go func(currencyTicker string) {
			getCurrencyData(currencyTicker)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}

// new goroutine
func getCurrencyData(currency string) {
	rate, err := api.GetRate("BTC")
	if err == nil {
		fmt.Printf("Rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}
