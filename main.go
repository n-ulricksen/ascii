package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
)

var columns int

var (
	columnHeader    = "Dec\tHex\tOct\tChar\t"
	headerSeparator = "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\t"
	colorReset      = "\033[0m"
	colorCyan       = "\033[36m"
)

func main() {
	// Parse command line flags
	flag.IntVar(&columns, "c", 3, "number of columns to print ASCII table")
	flag.Parse()

	if columns < 1 || columns > 5 {
		log.Fatal("Invalid # of columns, use 1-5.")
	}

	var buf bytes.Buffer
	var rows = 128/columns + 1

	buf.WriteString(getColumnHeaders(columnHeader, columns))
	buf.WriteString(getColumnHeaders(headerSeparator, columns))
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
	buf.WriteByte('\n')

	fmt.Print(buf.String())
}

func getColumnHeaders(text string, columns int) string {
	var buf bytes.Buffer

	buf.WriteByte('\n')
	for i := 0; i < columns; i++ {
		buf.WriteString(text)
		if i != columns-1 {
			buf.WriteByte('|')
			buf.WriteByte(' ')
		}
	}

	return buf.String()
}
