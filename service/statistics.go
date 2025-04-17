package service

import (
	"math"
	"sync"
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

func ProcessData() {
	var wg sync.WaitGroup
	data := make([]int, 0) // REVIEW: 共享切片未加锁
	var mu sync.Mutex      // 添加互斥锁

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) { // 传递 i 避免闭包问题
			defer wg.Done()
			mu.Lock() // 加锁
			data = append(data, i)
			mu.Unlock() // 解锁
		}(i)
	}
	wg.Wait()
}
