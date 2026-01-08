package main

import (
	"math"
	"fmt"
)

func main() {
	const inflationRate = 3.8
	var investmentAmount, years float64 = 1000, 10
	expectedReturnRate := 5.5
	
	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
