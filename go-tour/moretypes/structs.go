package moretypes

import "fmt"

type Vertex struct {
	X, Y int
}

//Testes com struct, com iniciação e alteração de valores via ponteiro
func PrintStruct() {
	fmt.Println("Print Struct...")
	v := Vertex{X: 1, Y: 3} //Named parameter é opcional
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

/*
Uma forma interessante de iniciar um struct genérico e instanciar ele em um slices:
s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
*/
