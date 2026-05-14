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
		l := evel(i, tokens)
		*i++
		n := evel(i, tokens)
		if tokens[*i] == "+" {
			return l + n
		} else if tokens[*i] == "-" {
			return l - n
		} else if tokens[*i] == "*" {
			return l * n
		}
	}
	if tokens[*i] != "(" && tokens[*i] != ")" {
		k, _ := strconv.Atoi(tokens[*i])
		return k
	}
}

func Polish(src string) {
	//region Scanner
	var s scanner.Scanner
	s.Init(strings.NewReader(src))

	var tokens []string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, s.TokenText())
	}
	//endregion
	return evel(*0, tokens)

}
