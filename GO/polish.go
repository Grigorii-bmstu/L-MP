package main

import (
	"strconv"
	"strings"
	"text/scanner"
)

func evel(i *int, tokens []string) int {
	if tokens[*i] == ")" {
		*i++
	}
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
}

func Polish(src string) int {
	//region Scanner
	var s scanner.Scanner
	s.Init(strings.NewReader(src))

	var tokens []string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, s.TokenText())
	}
	//endregion
	idx := 0
	return evel(&idx, tokens)

}
