package main

import (
	"bytes"
	"fmt"
)

var columns = 3
var colorReset = "\033[0m"
var colorCyan = "\033[36m"

func main() {
	var buf bytes.Buffer
	var rows = 128/columns + 1

	buf.WriteByte('\n')
	for i := 0; i < columns; i++ {
		buf.WriteString("Dec\tHex\tOct\tChar\t")
		if i != columns-1 {
			buf.WriteByte('|')
			buf.WriteByte(' ')
		}
	}
	buf.WriteByte('\n')
	for i := 0; i < columns; i++ {
		buf.WriteString("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\t")
		if i != columns-1 {
			buf.WriteByte('|')
			buf.WriteByte(' ')
		}
	}
	buf.WriteByte('\n')

	for row := 0; row < rows; row++ {
		for col := 0; col < columns; col++ {
			n := col*(rows) + row
			if n >= 128 {
				break
			}

			buf.WriteString(fmt.Sprintf("%d\t%#x\t%#o\t", n, n, n))
			if n > 32 {
				buf.WriteString(fmt.Sprintf("%s%c%s\t", colorCyan, n, colorReset))
			} else {
				buf.WriteByte('\t')
			}

			if col != columns-1 {
				buf.WriteByte('|')
				buf.WriteByte(' ')
			}
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}
