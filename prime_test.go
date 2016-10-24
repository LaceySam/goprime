package prime

import (
	"math/big"
	"testing"
)

var (
	testP   = NewPrimeGenerator(0, nil)
	max2047 = new(big.Int)
	max2048 = new(big.Int)
)

func init() {
	max2047.Exp(big.NewInt(2), big.NewInt(2047), nil)
	max2048.Exp(big.NewInt(2), big.NewInt(2048), nil)
}

func TestGenerateRandomNumberWithinBitRange(t *testing.T) {
	n, ok := testP.GenerateRandomNumber()
	if !ok {
		t.Fatal("Error encountered when generating random number")
	}
	lowerBound := n.Cmp(max2047)
	if lowerBound == -1 {
		t.Fatal("Random number less than lower bound")
	}

	upperBound := n.Cmp(max2048)
	if upperBound == 1 {
		t.Fatal("Random number greater than lower bound")
	}
}
