package handler

import "github.com/gin-gonic/gin"

type SolveRequest struct {
	IntegralId int      `json:"integralId" binding:"required"`
	Method     string   `json:"method" binding:"required"`
	Tolerance  *float64 `json:"tolerance" binding:"required"`

	A *float64 `json:"a,omitempty"`
	B *float64 `json:"b,omitempty"`
}

type SolveResponse struct {
	Value      float64 `json:"value"`
	Partitions int     `json:"partitions"`
}

func Solve() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
