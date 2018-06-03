package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("-143.234"))
}

func comma(s string) string {
	var buf bytes.Buffer
	mantissaStart := 0

	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		mantissaStart = 1
	}
	mantissaEnd := strings.Index(s, ".")
	if mantissaEnd == -1 {
		mantissaEnd = len(s)
	}
	mantissa := s[mantissaStart:mantissaEnd]
	fmt.Println(2 % 3)
	fmt.Println(len(mantissa))

	n := len(mantissa) % 3
	if n > 0 {
		buf.Write([]byte(mantissa[:n]))
		if len(mantissa) > n {
			buf.WriteString(",")
		}
	}

	for i, c := range mantissa[n:] {
		if i%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(c)
	}
	buf.WriteString(s[mantissaEnd:])

	return buf.String()
}
