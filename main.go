package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 1.23, 1}

	for _, taxRate := range taxRates {
		fm := filemanager.New(
			"./resources/prices.txt",
			fmt.Sprintf("./resources/taxRate_%.0f.json", taxRate*100),
		)
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fm)
		priceJob.Process()
	}

}
