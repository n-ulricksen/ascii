package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
)

// Command line flag
var columns int

var first32Chars = []string{
	"NUL", "SOH", "STX", "ETX", "EOT", "ENQ", "ACK", "BEL", "BS", "TAB", "LF", "VT", "FF", "CR", "SO", "SI", "DLE", "DC1", "DC2", "DC3", "DC4", "NAK", "SYN", "ETB", "CAN", "EM", "SUB", "ESC", "FS", "GS", "RS", "US", "Space",
}

var char127 = "DEL"

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
			if n == 127 {
				buf.WriteString(char127 + "\t")
			} else if n > 32 {
				buf.WriteString(fmt.Sprintf("%s%c%s\t", colorCyan, n, colorReset))
			} else {
				buf.WriteString(fmt.Sprintf("%s\t", first32Chars[n]))
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
