package methods

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// method is just a function with a receiver argument.
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// O ponteiro faz alterar o conteúdo da própria entidade que o invocou
// Passagem sem o ponteiro é feito via cópia, alterando somente o valor da cópia
// Neste caso, não passar o ponteiro, não iria alterar os valores de X e Y
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//A adição de uma função a um método só pode ser feita para structs no mesmo package
//O exemplo da função anterior, caso necessário ser feito em outro package, poderia
//ser feito como a forma abaixo. A chamada seria Scale(&v, 10)
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func MethodsFunction() {
	fmt.Println("Methods functions")

	//Declarando a interface
	var a Abser
	v := Vertex{X: 1.1, Y: 2.2}
	//Atribuindo a implementação da interface
	a = &v
	fmt.Println(a.Abs())
	fmt.Println(Abs(v))
	fmt.Println("Value before scale: ", v)
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	v.Scale(10)
	fmt.Println("Value after scale: ", v)
}
