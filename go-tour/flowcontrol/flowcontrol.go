package flowcontrol

import "fmt"

func ForFunction() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("The total sum is: %d\n", sum)
}
