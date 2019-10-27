package exercise

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	msg := fmt.Sprintf("%g: it doesn't support complex numbers", e)
	return msg
}

func SqrtWithE(f float64) (float64, error) {
	if f >= 0 {
		return math.Sqrt(f), nil
	}
	return 0, ErrNegativeSqrt(f)
}
