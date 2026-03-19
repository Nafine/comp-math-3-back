package handler

import (
	"comp-math-3/internal/algo"
	"comp-math-3/internal/numeric"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SolveRequest struct {
	FunctionId *int     `json:"functionId" binding:"required"`
	Method     string   `json:"method" binding:"required"`
	Tolerance  *float64 `json:"tolerance" binding:"required"`

	A *float64 `json:"a" binding:"required"`
	B *float64 `json:"b" binding:"required"`
}

type SolveResponse struct {
	Value      float64 `json:"value"`
	Partitions int     `json:"partitions"`
}

func Solve() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SolveRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
			return
		}

		ig := numeric.Integral{
			F:         numeric.GetFunction(*req.FunctionId),
			Tolerance: *req.Tolerance,
			N:         4,
			A:         *req.A,
			B:         *req.B,
		}

		solution, err := algo.Solve(req.Method, ig)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, SolveResponse{
			Value:      solution.Value,
			Partitions: solution.Partitions,
		})
	}
}
