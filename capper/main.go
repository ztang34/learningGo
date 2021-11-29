package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buf bytes.Buffer

	c := Capper{wtr: &buf}
	n, err := fmt.Fprintf(&c, "abcd")

	if err != nil {
		fmt.Println(fmt.Errorf("cannot convert to upper case"))
	} else {
		s := buf.String()
		fmt.Printf("%d of chars have been written to buffer: %s", n, s)
	}

}

type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (int, error) {
	upper := bytes.ToUpper(p)
	return c.wtr.Write(upper)
}
