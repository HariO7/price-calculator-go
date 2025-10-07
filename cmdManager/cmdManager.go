package cmdmanager

import "fmt"

type CMDManger struct {
}

func (cm CMDManger) ReadDataFromFile() ([]string, error) {
	fmt.Println("Please enter your Prices, Confirm every price with enter: ")

	var prices []string

	for {
		var price string
		fmt.Println("Enter Price:")
		fmt.Scan(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cm CMDManger) WriteJSON(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManger {
	return CMDManger{}
}
