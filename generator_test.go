package fibonacci

import (
	"math"
	"strconv"
	"testing"
)

func TestFibonacciNegative(t *testing.T) {
	var test_values = []int{-1, -2, -3, -10, -1000, -1000000, math.MinInt64}
	for _, i := range test_values {
		if _, err := NewGenerator(i); err == nil {
			t.Errorf("Expected NewGenerator to return error when asked for %d iterations", i)

		}
	}
}

func TestFibonacciNumberOfValuesGenerated(t *testing.T) {
	var test_values = []int{1, 2, 3, 10, 20, 100, 1000, 100000}

	for _, i := range test_values {
		fg, err := NewGenerator(i)
		if err != nil {
			t.Errorf("Expected NewGenerator to return error when asked for %d iterations", i)
		}

		result_chan := make(chan FibNum)
		go fg.Produce(result_chan)

		cnt := 0
		for dont_care := range result_chan {
			_ = dont_care
			cnt = cnt + 1
		}

		if cnt != i {
			t.Errorf("Generated %d numbers when asked to generate %d, maxIterations is %d", cnt, i, fg.maxIterations)
		}
	}
}

func TestFibonacciVerifyCorrectOutputAgainstInt(t *testing.T) {
	//At 'iteration' 93 int rolls and cannot be used to validate the implemented algorithm
	var test_values = []int{1, 2, 3, 10, 20, 93}

	for _, i := range test_values {
		fg, err := NewGenerator(i)
		if err != nil {
			t.Errorf("Expected NewGenerator to return error when asked for %d iterations", i)
		}

		result_chan := make(chan FibNum)
		go fg.Produce(result_chan)

		cnt := 0
		x, y := 0, 1
		for val := range result_chan {
			cnt = cnt + 1
			if x < 0 {
				t.Fatalf("Test problem on %d iteration: expected value went negative %d", cnt, x)
			}
			if strconv.Itoa(x) != val.String() {
				t.Fatalf("Incorrect value on iteration %d of test %d: expected %d got %s",
					cnt, i, x, val.String())
			}
			x, y = y, x+y
		}
	}
}
