package prime

import (
	"math/big"
	"math/rand"
	"time"
)

const (
	minimumBits          = 2048
	minimumPrimeAccuracy = 100
)

// PrimeGenerator generates random large primes
type PrimeGenerator struct {
	Bits              int
	PrimeAccuracy     int
	Source            rand.Source
	randomGenerator   *rand.Rand
	preComputedPrimes []*big.Int
}

// GetPreComputedPrimes returns an array of the first 1000 primes as []*big.Int
func GetPreComputedPrimes() []*big.Int {
	biggedPreComputedPrimes := []*big.Int{}
	for _, prime := range preComputedPrimes {
		biggedPreComputedPrimes = append(biggedPreComputedPrimes, big.NewInt(prime))
	}

	return biggedPreComputedPrimes
}

// NewPrimeGenerator returns a PrimeGenerator with atleast the default fields
func NewPrimeGenerator(bits, primeAccuracy int, source rand.Source) *PrimeGenerator {
	if bits == 0 || bits < minimumBits {
		bits = minimumBits
	}

	if primeAccuracy == 0 || primeAccuracy < minimumPrimeAccuracy {
		primeAccuracy = minimumPrimeAccuracy
	}

	if source == nil {
		// TODO(Sam): Investiagte better randon source?
		source = rand.NewSource(time.Now().UnixNano())
	}

	biggedPreComputedPrimes := GetPreComputedPrimes()

	return &PrimeGenerator{
		Bits:              bits,
		PrimeAccuracy:     primeAccuracy,
		Source:            source,
		randomGenerator:   rand.New(source),
		preComputedPrimes: biggedPreComputedPrimes,
	}
}

// GetPrime generates a new large prime of n bits
func (p *PrimeGenerator) GetPrime() (*big.Int, bool) {

	// We're going to need to keep generating random numbers until we prove one is a prime.
	// Should get a prime within a couple of hundred attempts.
	prime, ok := p.GenerateRandomNumber()
	if !ok {
		return &big.Int{}, ok
	}

	for !p.ProvePrimality(prime) {
		prime, ok = p.GenerateRandomNumber()
		if !ok {
			return &big.Int{}, ok
		}
	}

	return prime, true
}

// GenerateRandomNumber generates a random number of n bits
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

// checkAgainstSmallPrimes checks the mod of the candidate against the first 1k primes
// In many cases this weeds out easily proven non primes before the more expensive test
func (p *PrimeGenerator) checkAgainstSmallPrimes(candidate *big.Int) bool {
	zero := big.NewInt(0)
	mod := new(big.Int)

	for _, prime := range p.preComputedPrimes {
		if mod.Mod(candidate, prime) == zero {
			return false
		}
	}

	return true
}

// ProvePrimality will attempt to prove if a given number is a prime. It has a level of innacuracy
func (p *PrimeGenerator) ProvePrimality(candidate *big.Int) bool {
	if !p.checkAgainstSmallPrimes(candidate) {
		return false
	}

	return candidate.ProbablyPrime(p.PrimeAccuracy)
}
