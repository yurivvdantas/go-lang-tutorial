package main

import (
	"fmt"
	"go-tour/flowcontrol"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

//Exemplo de variáveis sendo declaradas em uma única linha
//O tipo vem como último parâmetro
var d, python, java bool = false, true, false

//Forma mais elegante de iniciar as variáveis
var (
	ToBe              = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//Exemplo de retorno sendo ignorado, no caso o error
func add(x, y int) (int, error) {
	return x + y, nil
}

//naked return. Retorna os mesmos parâmetros no retorno
//Só deve ser usado em funções curtas
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func randAndMath() {
	rand.Seed(time.Now().Unix())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Printf("The value of pi is %g \n", math.Pi)
}

func zeroValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println("Print default values:")
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func testConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func main() {

	a, b, c := "a", false, 1
	fmt.Println(a, b, c, c, d, python, java)
	fmt.Println(ToBe, MaxInt, z)
	randAndMath()
	i, _ := add(1, 3)
	fmt.Printf("Sum is %v\n", i)
	fmt.Println(split(17))
	zeroValues()
	testConversion()
	flowcontrol.ForFunction()
	TestImportFunctionFromSamePackage()
}
