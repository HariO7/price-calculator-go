package main

import (
	cmdmanager "example.com/price-calculator/cmdManager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 1.23, 1}

	for _, taxRate := range taxRates {
		// fm := filemanager.New(
		// 	"./resources/prices.txt",
		// 	fmt.Sprintf("./resources/taxRate_%.0f.json", taxRate*100),
		// )
		cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, cmd)
		priceJob.Process()
	}

}
