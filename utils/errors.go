package utils

import "errors"

func Divider(numerator float64, denominator float64) (float64, error) {
	if denominator == 0 {
		return 0, errors.New("denominator is 0")
	}

	return numerator / denominator, nil
}
