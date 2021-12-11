package methods

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//Testes com struct, com iniciação e alteração de valores via ponteiro
func PrintStruct() {
	fmt.Println("Print Struct...")
	v := Vertex{X: 1, Y: 3} //Named parameter é opcional
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
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
