package flowcontrol

import (
	"fmt"
	"math"
	"runtime"
)

func ForFunction() {
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
}

func IfFunction(x, n, lim float64) float64 {
	//variável v fica disponível apenas no escopo do bloco if ou do else vínculado
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func SwitchFunction() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
