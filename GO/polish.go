package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/scanner"
)

func evel(i *int, tokens []string) int {

	if tokens[*i] == "(" {
		*i++
		v := tokens[*i]
		*i++
		l := evel(i, tokens)
		n := evel(i, tokens)
		*i++
		if v == "+" {
			return l + n
		} else if v == "-" {
			return l - n
		} else if v == "*" {
			return l * n
		}
	} else {
		k, _ := strconv.Atoi(tokens[*i])
		*i++
		return k
	}
	return 0
}

func Polish(src string) int {
	//region Scanner
	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Mode = scanner.ScanIdents | scanner.ScanStrings | scanner.ScanComments
	var tokens []string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, s.TokenText())
	}
	//endregion
	idx := 0
	return evel(&idx, tokens)

}
func main() {

	sc := bufio.NewScanner(os.Stdin)
	var fullInput strings.Builder

	for sc.Scan() {
		fullInput.WriteString(sc.Text())
		fullInput.WriteString(" ")
	}

	expr := strings.TrimSpace(fullInput.String())
	if expr != "" {
		fmt.Println(Polish(expr))
	}
}
