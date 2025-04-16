package handlers

import (
	// "math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VarianceRequest struct {
	Numbers []float64 `json:"numbers"`
}

type VarianceResponse struct {
	Variances []float64 `json:"variances"`
	Error     string    `json:"error,omitempty"`
}

func (h *Handler) CalculateVariance(c *gin.Context) {
	var req VarianceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, VarianceResponse{Error: "Invalid JSON payload"})
		return
	}

	if len(req.Numbers) == 0 {
		c.JSON(http.StatusBadRequest, VarianceResponse{Error: "Numbers array cannot be empty"})
		return
	}

	// Calculate mean
	sum := 0.0
	for _, num := range req.Numbers {
		sum += num
	}
	mean := sum / float64(len(req.Numbers))

	// Calculate variance
	variance := 0.0
	for _, num := range req.Numbers {
		diff := num - mean
		variance += diff * diff
	}
	variance = variance / float64(len(req.Numbers))

	c.JSON(http.StatusOK, VarianceResponse{Variances: []float64{variance}})
}
