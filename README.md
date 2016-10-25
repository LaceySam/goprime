# goprime
Go library to generate large prime numbers 2048 bits in size or more.

Note:

- Can take a while to find a prime (I've seen it take anywhere between 1s - 20s)
- Does not guarentee a legitimate prime, only as high probablility that it's
  found one, see:

  https://en.wikipedia.org/wiki/Miller%E2%80%93Rabin_primality_test

Benchmarks:

BenchmarkGetPrime-4                100   5599447153 ns/op 2949377011 B/op 3572507 allocs/op  
BenchmarkGenerateRandomNumber-4    1000  1738294 ns/op    2243738 B/op    2060 allocs/op  
BenchmarkCheckAgainstSmallPrimes-4 10000 167983 ns/op     48080 B/op      1002 allocs/op  
BenchmarkPrimailityTest-4          2000  749293 ns/op     68145 B/op      1309 allocs/op  
