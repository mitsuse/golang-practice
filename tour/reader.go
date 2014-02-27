package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (this *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = this.r.Read(p)
	for i := 0; i < n; i++ {
		if ('a' <= p[i] && p[i] <= 'z'-13) || ('A' <= p[i] && p[i] <= 'Z'-13) {
			p[i] += 13
		} else if ('z'-13 < p[i] && p[i] <= 'z') || ('Z'-13 < p[i] && p[i] <= 'Z') {
			p[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
