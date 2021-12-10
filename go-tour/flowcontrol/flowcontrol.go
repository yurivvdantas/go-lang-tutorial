package flowcontrol

import (
	"fmt"
	"math"
	"time"
)

func ForFunction() {
	fmt.Println("For function...")
	sum := 0
	for i := 0; i <= 10; i++ {
		if sum%2 == 0 {
			fmt.Printf("%d, ", sum)
		}
		sum += i
	}
	fmt.Printf("The total sum is: %d\n", sum)
	/*
			O while é feito com for sem os parâmetros de entrada e de incremento:
			for sum < 1000 {
			   sum += sum
		    }
	*/
	//Também é possível usar o range para o for
	//When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
	//Também é possível usar das seguintes formas:
	//for i, _ := range pow, for _, v := range pow, for i := range pow
	for i, v := range []int{10, 20, 30, 40} {
		fmt.Printf("i: %d, v: %d\n", i, v)
	}
}

func IfFunction(x, n, lim float64) float64 {
	//variável v fica disponível apenas no escopo do bloco if ou do else vínculado
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func SwitchFunction() {
	fmt.Println("Switch Function...")
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	/*
			Switch sem uma condição é o mesmo que true:
		t := time.Now()
		switch {
			case t.Hour() < 12:
				fmt.Println("Good morning!")
			case t.Hour() < 17:
				fmt.Println("Good afternoon.")
			default:
				fmt.Println("Good evening.")
		}
	*/
}

// A defer statement defers the execution of a function until the surrounding function returns.
//The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
func DeferFunction() {
	fmt.Println("Defer function...")
	fmt.Println("counting")

	for i := 0; i < 2; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
