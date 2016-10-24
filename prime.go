package prime

import (
	"math/big"
	"math/rand"
	"time"
)

const (
	minimumBits = 2048
)

type PrimeGenerator struct {
	Bits            int
	Source          rand.Source
	randomGenerator *rand.Rand
}

func NewPrimeGenerator(bits int, source rand.Source) *PrimeGenerator {
	if bits == 0 || bits < minimumBits {
		bits = minimumBits
	}

	if source == nil {
		// TODO(Sam) Investiagte better randon source?
		source = rand.NewSource(time.Now().UnixNano())
	}

	return &PrimeGenerator{
		Bits:            bits,
		Source:          source,
		randomGenerator: rand.New(source),
	}
}

func (p *PrimeGenerator) GetPrime() (*big.Int, bool) {

	// We're going to need to keep generating random numbers until we prove one is a prime.
	// Should get a prime within a couple of hundred attempts.

	prime, ok := p.GenerateRandomNumber()
	if !ok {
		return &big.Int{}, ok
	}

	for p.ProvePrime(prime) {
		prime, ok = p.GenerateRandomNumber()
		if !ok {
			return &big.Int{}, ok
		}
	}

	return prime, true
}

func (p *PrimeGenerator) GenerateRandomNumber() (*big.Int, bool) {
	// We're going to go through a binary sequence and randomly turn the bits on and off.
	// First bit is always on to ensure we get a random number of: 2^(b-1) <= x <= 2^b where b is
	// bytes.

	binarySequence := "1"
	for i := 0; i < p.Bits-1; i++ {

		if p.randomGenerator.Intn(2) == 0 {
			binarySequence += "0"
			continue
		}

		binarySequence += "1"
	}

	randomNumber := new(big.Int)
	return randomNumber.SetString(binarySequence, 2)
}

func (p *PrimeGenerator) ProvePrime(prime *big.Int) bool {
	return true
}
