package rounding

import (
	"fmt"
	"testing"
)

func TestRoundingFloat(t *testing.T) {
	amount := float64(10)
	result := RoundFloat(amount, 2)
	fmt.Println("Amount: ", result)
}
