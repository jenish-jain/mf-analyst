package util

import (
	"fmt"
	"strconv"
)

func GetFloat64FromString(amount string) float64 {
	floatAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Printf("error parsing string amount to float %s amount - %v \n", amount, err)
		//TODO check either to panic to return default amount here
	}
	return floatAmount
}
