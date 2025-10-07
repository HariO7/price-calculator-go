package main

import (
	"fmt"

	cmdmanager "example.com/price-calculator/cmdManager"
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 1.23, 1}

	var choice string

	fmt.Println("Select your mode \n 1.File written \n 2.command line")

	fmt.Scan(&choice)

	for _, taxRate := range taxRates {
		if choice == "1" {
			fm := filemanager.New(
				"./resources/prices.txt",
				fmt.Sprintf("./resources/taxRate_%.0f.json", taxRate*100),
			)
			priceJob := prices.NewTaxIncludedPriceJob(taxRate, fm)
			priceJob.Process()
		} else {
			cmd := cmdmanager.New()
			priceJob := prices.NewTaxIncludedPriceJob(taxRate, cmd)
			priceJob.Process()
		}

	}

}
