package service

import (
	"math"
)

// CalculateStandardDeviation computes the standard deviation of a given array of numbers.
func CalculateStandardDeviation(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	// Calculate the mean
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	mean := sum / float64(len(numbers))

	// Calculate the variance
	var varianceSum float64
	for _, num := range numbers {
		varianceSum += math.Pow(num-mean, 2)
	}
	variance := varianceSum / float64(len(numbers))

	// Return the square root of the variance (standard deviation)
	return math.Sqrt(variance)
}
