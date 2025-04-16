package services

import (
	"errors"
	"math"
)

func (s *Service) CalculateVariance(numbers []float64) ([]float64, error) {
	if len(numbers) == 0 {
		return nil, errors.New("input array cannot be empty")
	}

	// Calculate the mean
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	mean := sum / float64(len(numbers))

	// Calculate variance for each element
	variances := make([]float64, len(numbers))
	for i, num := range numbers {
		variances[i] = math.Pow(num-mean, 2)
	}

	return variances, nil
}
