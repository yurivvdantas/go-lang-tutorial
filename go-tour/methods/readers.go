package methods

import (
	"fmt"
	"io"
	"strings"
)

//Read populates the given byte slice with data and returns the number of bytes populated
//and an error value. It returns an io.EOF error when the stream ends.
func ReaderFunction() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
