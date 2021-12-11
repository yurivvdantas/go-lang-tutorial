package moretypes

import (
	"fmt"
	"go-tour/methods"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	listS := strings.Fields(s)
	mapString := make(map[string]int)
	for _, v := range listS {
		count, ok := mapString[v]
		if !ok {
			mapString[v] = 1
		} else {
			mapString[v] += count
		}
	}
	return mapString
}

func MapFunction() {
	fmt.Println("Map function...")
	m := make(map[string]methods.Vertex)
	m["Bell Labs"] = methods.Vertex{
		X: 40, Y: -74,
	}

	m = map[string]methods.Vertex{
		"Bell Labs": {
			X: 41, Y: -74,
		},
		"Google": {
			X: 37, Y: -122,
		},
	}
	m["microsoft"] = methods.Vertex{X: 3, Y: 4}
	delete(m, "Google")

	//If key is in m, ok is true. If not, ok is false.
	//If key is not in the map, then elem is the zero value for the map's element type.
	elem, ok := m["Google"]

	//var m2 map[string]Vertex
	//copy(m2, m)

	fmt.Println("value of m: ", m["Bell Labs"], ". Value of elem: ", elem, ". Elem is present? ", ok)

	wc.Test(WordCount)
}
