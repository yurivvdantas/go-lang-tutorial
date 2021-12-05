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
