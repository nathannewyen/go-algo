package mathalgo

import "testing"

func TestGCD(t *testing.T) {
	result := GreatestCommonDivisor(48, 18)
	if result != 6 {
		t.Errorf("expected GCD 6, got %d", result)
	}
}

func TestLCM(t *testing.T) {
	result := LeastCommonMultiple(12, 18)
	if result != 36 {
		t.Errorf("expected LCM 36, got %d", result)
	}
}

func TestFastPower(t *testing.T) {
	result := FastPower(2, 10)
	if result != 1024 {
		t.Errorf("expected 1024, got %d", result)
	}
}

func TestIsPrime(t *testing.T) {
	if !IsPrime(17) {
		t.Error("17 should be prime")
	}
	if IsPrime(15) {
		t.Error("15 should not be prime")
	}
}

func TestSieve(t *testing.T) {
	primes := SieveOfEratosthenes(20)
	if len(primes) != 8 {
		t.Errorf("expected 8 primes up to 20, got %d", len(primes))
	}
}

func TestFibonacci(t *testing.T) {
	if FibonacciIterative(10) != 55 {
		t.Errorf("expected fib(10)=55, got %d", FibonacciIterative(10))
	}
	if FibonacciMemoized(10) != 55 {
		t.Errorf("expected fib(10)=55, got %d", FibonacciMemoized(10))
	}
}
