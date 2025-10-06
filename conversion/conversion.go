package conversion

import (
	"errors"
	"strconv"
)

func ConvertStringsTofloat(strings []string) ([]float64, error) {

	convertedVal := make([]float64, len(strings))

	for strIdx, str := range strings {
		floatPrice, err := strconv.ParseFloat(str, 64)

		if err != nil {
			return nil, errors.New("error occured while converting strings to floats")
		}

		convertedVal[strIdx] = floatPrice
	}

	return convertedVal, nil
}
