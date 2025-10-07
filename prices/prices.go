package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64              `json:"tax_rate"`
	InputPrices       []float64            `json:"input_prices"`
	TaxIncludedPrices map[string]string    `json:"tax_included_prices"`
	IOManger          filemanager.IOManger `json:"-"`
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	job.IOManger.WriteJSON(job)

}

func NewTaxIncludedPriceJob(taxRate float64, iom filemanager.IOManger) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
		IOManger:    iom,
	}
}

func (job TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManger.ReadDataFromFile()

	if err != nil {
		return err
	}

	prices, err := conversion.ConvertStringsTofloat(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}
