// echo implementted in golang
// author: synkro
// version: 0.1

package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ReplaceBackslashEscapes(s string) string {
	r := strings.NewReplacer(
		`\\`, "\\",
		`\a`, "\a",
		`\n`, "\n",
		`\b`, "\b",
		`\e`, "\x1b",
		`\f`, "\f",
		`\r`, "\r",
		`\t`, "\t",
		`\v`, "\v",
	)
	return r.Replace(s)
}

// ToAscii converts the hex text format \xNN to a string
func ToAscii(s string) string {
	// dont forget to  handle this error but not for now dw
	hex := s[2:] // remove the leading `\x`
	value, err := strconv.ParseInt(hex, 16, 16)
	if err != nil {
		panic(err)
	}
	return string(rune(value))
}

func ConvertHexValues(s string) string {
	r := regexp.MustCompile(`\\x[0-9a-zA-Z]{2}`) // dont forget the ` ` with regex to make it a literal
	res := r.ReplaceAllStringFunc(s, ToAscii)
	return res
}

func main() {
	newLine := flag.Bool("n", false, "do not output trailing whiteline")
	enableBackslash := flag.Bool("e", false,
		`enable interpretation of backslash escapes
	\\ backslash
	\a alert (BEL)
	\b backspace
	\e escape 
	\f form feed 
	\n new line
	\r carriage return
	\t horizontal tab
	\v vertical tab
	`)
	flag.Parse()

	args := flag.Args()
	out := strings.Join(args, " ")

	if *enableBackslash {
		out = ReplaceBackslashEscapes(out)
		out = ConvertHexValues(out)
	}

	if !*newLine {
		fmt.Println(out)
	} else {
		fmt.Print(out)
	}
}
