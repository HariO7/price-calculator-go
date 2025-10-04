package main

import (
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 1.23, 1}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPrice(taxRate)
		priceJob.Process()
	}

}
