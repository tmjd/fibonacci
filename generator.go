package fibonacci

import (
	"errors"
	"math/big"
)

// Custom Number to use generating the fibonacci numbers, can swap out
// the underlying type without affecting the algorithm
type FibNum struct {
	value big.Int
}

func newFibNum(init int64) FibNum {
	fn := FibNum{}
	fn.value = *big.NewInt(init)
	return fn
}

// Clone to get copy of a FibNum.
// With big.Int need to do a new big.NewInt so the FibNum is not using the same
// big.Int value.
func cloneFibNum(src FibNum) FibNum {
	fn := FibNum{}
	fn.value = *big.NewInt(0)
	fn.value.Set(&src.value)
	return fn
}

func (fn *FibNum) add(a FibNum, b FibNum) {
	fn.value.Add(&a.value, &b.value)
}

// Return string representation of FibNum
func (fn FibNum) String() string {
	return fn.value.String()
}

type Generator struct {
	maxIterations int
}

func NewGenerator(iterations int) (fg *Generator, err error) {
	if iterations < 0 {
		return nil, errors.New("Number of iterations cannot be negative")
	} else if iterations > 100000 {
		// Seems like an unreasonably high number but I think there should be some
		// limit to what will be generated
		return nil, errors.New("Number of iterations cannot be greater than 100000")
	}
	fg = &Generator{}
	fg.maxIterations = iterations
	return fg, nil
}

func (fg *Generator) Produce(out chan<- FibNum) {
	if fg.maxIterations == 0 {
		close(out)
		return
	}
	var v [2]FibNum
	v[0] = newFibNum(0)
	v[1] = newFibNum(1)
	idx := 0

	for i := 0; i < fg.maxIterations; i = i + 1 {
		out <- cloneFibNum(v[idx])
		v[idx].add(v[0], v[1])
		if idx == 0 {
			idx = 1
		} else {
			idx = 0
		}
	}

	close(out)
}
