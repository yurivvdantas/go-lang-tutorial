package methods

import (
	"fmt"
	"strings"
)

type I interface {
	M()
	F()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

//Caso a interface possua mais de um método, o compilador obriga
//a implementação de todos os métodos
func (t T) F() {
	fmt.Println(t.S)
}

func InterfaceFunction() {
	fmt.Print("Interface function...")
	var i I = &T{"hello"}
	i.M()

	var t *T
	i = t
	//i possui valor nil, mas ao chamar o método M não é lançado nullPointer
	//isso permite o tratamento dentro do próprio método
	i.M()
	genericInterface()
	checkTypeInterface()
}

func genericInterface() {
	//An empty interface may hold values of any type.
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
	exerciseStringers()
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

//Sempre usar t, ok := i.(T) para tratar erro e evitar panic
func checkTypeInterface() {
	fmt.Print("Check interface...")
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)
	fmt.Println(s)

	f, ok := i.(float64)
	fmt.Println(f, ok)
	fmt.Println(s)

	//f = i.(float64)
	//fmt.Println(f) panic
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	s := ""
	for _, v := range ip {
		//s = s + string
		s = s + fmt.Sprintf("%v.", v)
	}
	return strings.TrimRight(s, ".")
}

func exerciseStringers() {
	//https://go.dev/tour/methods/18
	fmt.Println("Exercise stringers ip...")
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

//É possível usar o switch com type:
/*
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
*/

//É possível sobrescrever a implementação do fmt.print implementando a interface String()
/*
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
*/
//https://go.dev/tour/methods/17
