package primenum

import (
	"math"
	"slices"
	pkg_sort "sort"
)

func IsPrimeNumber(input int64) bool {
	if input <= 1 {
		return false
	}

	ok := CacheExists(input)
	if ok {
		return true
	}

	if isNotPrimNumInCache(input) {
		return false
	}

	i := int64(2)
	n := input
	if sqRoot := squareRoot(input); sqRoot > 2 {
		n = sqRoot
	}
	for i <= n {
		if input%i == 0 && input != i {
			return false
		}
		i += 1
	}
	return true
}

func IsPrimeNumberNoCache(input int64) bool {
	if input <= 1 {
		return false
	}
	i := int64(2)
	n := input
	for i < n {
		if n%i == 0 {
			return false
		}
		i += 1
	}
	return true
}

func isNotPrimNumInCache(input int64) bool {
	result := false
	CacheEach(func(n int64) bool {
		if input%n == 0 {
			result = true
			return false
		}
		return true
	})
	return result
}

// Eratosthenes ...
func Eratosthenes(input int64) []int64 {
	isNotPrimes := make([]bool, input+1)

	isNotPrimes[0], isNotPrimes[1] = true, true
	sqRoot := squareRoot(input)
	for i := int64(2); i <= sqRoot; i++ {
		if isNotPrimes[i] {
			continue
		}
		for j := i * 2; j <= input; j += i {
			isNotPrimes[j] = true
		}
	}

	result := make([]int64, 0, sqRoot)
	for i, isNotPrime := range isNotPrimes {
		if !isNotPrime {
			result = append(result, int64(i))
		}
	}
	updateCache(result)
	return result
}

func updateCache(s []int64) {
	m := GetMaxInCache()
	if m < 2 || s[0] > m {
		CacheStore(s...)
	} else if s[len(s)-1] > m {
		idx, ok := slices.BinarySearch(s, m)
		if !ok {
			panic("error")
		}
		CacheStore(s[idx+1:]...)
	}
}

func PrimeFactorization(input int64) []int64 {
	if input < 2 {
		return []int64{}
	}

	result := make([]int64, 0, input)

	initLoopIndex := func() int64 { return 2 }

	i := initLoopIndex()
	n := input
	for i <= n {
		if IsPrimeNumber(n) {
			result = append(result, n)
			break
		}
		if IsPrimeNumber(i) && n%i == 0 {
			result = append(result, i)
			n /= i
			i = initLoopIndex()
			continue
		}
		i += 1
	}
	return sort(result)
}

func sort(s []int64) []int64 {
	pkg_sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s
}

func PrimeNumbers(input int64) []int64 {
	if input < 2 {
		return []int64{}
	}
	resultSize := input
	if sqRoot := squareRoot(input); sqRoot > 2 {
		resultSize = sqRoot
	}
	result := make([]int64, 0, resultSize)

	for i := int64(2); i <= input; i++ {
		if IsPrimeNumber(i) {
			result = append(result, i)
		}
	}
	return result
}

func squareRoot(input int64) int64 {
	return int64(math.Floor(math.Sqrt(float64(input))))
}
