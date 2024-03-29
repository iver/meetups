package fib

import "testing"

// use values that you know are right
var tests = []uint64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

func TestFibFunc(t *testing.T) {
	fn := FibFunc()
	for i, v := range tests {
		if val := fn(); val != v {
			t.Fatalf("at index %d, expected %d, got %d.", i, v, val)
		}
	}
}

func TestOther(t *testing.T) {
	t.Log("Init ....")
	result := 4 + 9
	t.Logf("Result: %v", result)
}
