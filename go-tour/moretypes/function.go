package moretypes

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func FunctionsValues() {
	fmt.Println("Functions...")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	closureFunction()
	fibonacciFunction()
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

/*
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
For example, the adder function returns a closure. Each closure is bound to its own sum variable.
*/
func closureFunction() {
	pos, neg := adder(), adder()
	for i := 0; i < 3; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	lastItem := 1
	afterLastItem := 1
	cursor := 1
	return func() int {
		nextItem := lastItem + afterLastItem
		cursor = lastItem
		lastItem = afterLastItem
		afterLastItem = nextItem
		return cursor
	}
}

func fibonacciFunction() {
	fmt.Println("Fibonacci Function")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), ",")
	}
	fmt.Println()
}
