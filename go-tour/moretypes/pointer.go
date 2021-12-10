package moretypes

import "fmt"

func PointerFunction() {
	fmt.Println("Pointer function...")
	i := 42

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

}
