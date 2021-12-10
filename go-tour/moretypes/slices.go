package moretypes

import "fmt"

//Array possui um tamanho pré definido, enquanto slices não.
//Slices são muito mais comuns
//O retorno como ponteiro foi para permitir a instanciação por um slice
//no método sliceFunction. Mais informações aqui: https://stackoverflow.com/questions/50062243/error-addressing-the-returned-slice-of-a-function
func arrayFunction() *[4]string {
	//Poderia ter sido iniciado como [4]string{"a","b","c","d"}
	var a [4]string
	a[0] = "a"
	a[1] = "b"
	a[2] = "c"
	a[2] = "d"
	return &a
}

func SlicesFunction() {
	fmt.Println("Slices function")
	//This selects a half-open range which includes the first element, but excludes the last one.
	s := arrayFunction()[0:4]
	s[0] = "z"
	// When slicing, you may omit the high or low bounds to use their defaults instead.
	//The default is zero for the low bound and the length of the slice for the high bound.
	//Ex: arrayFunction()[0:], arrayFunction()[:4], arrayFunction()[:]
	b := arrayFunction()[0:4]
	c := b[0:2]
	//Changing the elements of a slice modifies the corresponding elements of its underlying array.
	c[0] = "x"
	fmt.Println(s)
	fmt.Println(b)
	sliceCapacityAndLength()
	makeSlice()
	appendSlice()
}

// The length of a slice is the number of elements it contains.
// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
func sliceCapacityAndLength() {
	fmt.Print("Slice capacity and length")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

//The make function allocates a zeroed array and returns a slice that refers to that array:
//To specify a capacity, pass a third argument to make
func makeSlice() {
	fmt.Print("make slice")
	a := make([]int, 5)
	printSlice(a)
	b := make([]int, 0, 5)
	printSlice(b)
	c := b[:2]
	printSlice(c)
	d := c[2:5]
	printSlice(d)
}

func appendSlice() {
	s := arrayFunction()[0:4]
	s = append(s, "1", "2", "3", "4")
	fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/*
Slices can contain any type, including other slices.
board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
*/
