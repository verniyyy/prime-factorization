package primenum

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func BenchmarkIsPrimeNumber(b *testing.B) {
	b.Run("cached", func(b *testing.B) {
		CacheClear()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			n := randNum()
			_ = IsPrimeNumber(n)
		}
	})
	b.Run("no cached", func(b *testing.B) {
		CacheClear()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			n := randNum()
			_ = IsPrimeNumberNoCache(n)
		}
	})
}

func BenchmarkPrimeNumbers(b *testing.B) {
	n := int64(99999)
	b.Run("caching", func(b *testing.B) {
		CacheClear()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = PrimeNumbers(n)
		}
	})
	b.Run("eratosthenes", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Eratosthenes(n)
		}
	})
}

func BenchmarkPrimeFactorization(b *testing.B) {
	n := int64(999)
	b.Run("default", func(b *testing.B) {
		CacheClear()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = PrimeFactorization(n)
		}
	})
	b.Run("use eratosthenes", func(b *testing.B) {
		CacheClear()
		_ = Eratosthenes(n)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = PrimeFactorization(n)
		}
	})
}

func randNum() int64 {
	return rand.Int63n(999999)
}

func ExamplePrimeFactorization() {
	n := int64(999999)
	result := PrimeFactorization(n)
	fmt.Printf("result: %v\n", result)
	// Output:
	// result: [3 3 3 7 11 13 37]
}

// ExamplePrimeNumbers ...
func ExamplePrimeNumbers() {
	n := int64(999)
	result := PrimeNumbers(n)
	fmt.Printf("result: %v\n", result)
	// Output:
	// result: [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97 101 103 107 109 113 127 131 137 139 149 151 157 163 167 173 179 181 191 193 197 199 211 223 227 229 233 239 241 251 257 263 269 271 277 281 283 293 307 311 313 317 331 337 347 349 353 359 367 373 379 383 389 397 401 409 419 421 431 433 439 443 449 457 461 463 467 479 487 491 499 503 509 521 523 541 547 557 563 569 571 577 587 593 599 601 607 613 617 619 631 641 643 647 653 659 661 673 677 683 691 701 709 719 727 733 739 743 751 757 761 769 773 787 797 809 811 821 823 827 829 839 853 857 859 863 877 881 883 887 907 911 919 929 937 941 947 953 967 971 977 983 991 997]
}

func TestPrimeFactorization(t *testing.T) {
	type args struct {
		input int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
		{
			name: "N/A by negative number",
			args: args{
				input: -1,
			},
			want: []int64{},
		},
		{
			name: "N/A by zero",
			args: args{
				input: 0,
			},
			want: []int64{},
		},
		{
			name: "N/A by 1",
			args: args{
				input: 1,
			},
			want: []int64{},
		},
		{
			name: "itself is a prime number of 2",
			args: args{
				input: 2,
			},
			want: []int64{2},
		},
		{
			name: "itself is a prime number of 3",
			args: args{
				input: 3,
			},
			want: []int64{3},
		},
		{
			name: "itself is a prime number of 5",
			args: args{
				input: 5,
			},
			want: []int64{5},
		},
		{
			name: "itself is a prime number of 7",
			args: args{
				input: 7,
			},
			want: []int64{7},
		},
		{
			name: "itself is a prime number of 11",
			args: args{
				input: 11,
			},
			want: []int64{11},
		},
		{
			name: "highly composite number",
			args: args{
				input: 735134400,
			},
			want: []int64{2, 2, 2, 2, 2, 2, 3, 3, 3, 5, 5, 7, 11, 13, 17},
		},
		{
			name: "product of two large prime numbers",
			args: args{
				input: 45577073,
			},
			want: []int64{4637, 9829},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrimeFactorization(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrimeFactorization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrimeNumber(t *testing.T) {
	type args struct {
		input int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "2",
			args: args{
				input: 2,
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				input: 3,
			},
			want: true,
		},
		{
			name: "5908",
			args: args{
				input: 5908,
			},
			want: false,
		},
		{
			name: "8",
			args: args{
				input: 8,
			},
			want: false,
		},
		{
			name: "1681",
			args: args{
				input: 1681,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrimeNumber(tt.args.input); got != tt.want {
				t.Errorf("IsPrimeNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEratosthenes(t *testing.T) {
	type args struct {
		input int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				input: 21,
			},
			want: []int64{2, 3, 5, 7, 11, 13, 17, 19},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Eratosthenes(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Eratosthenes() = %v, want %v", got, tt.want)
			}
		})
	}
}
