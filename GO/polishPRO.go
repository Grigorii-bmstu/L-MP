package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/scanner"
)

func evel(i *int, tokens []string, m map[string]int) string {
	if *i >= len(tokens) {
		return ""
	}

	if tokens[*i] == "(" {
		*i++
		if *i >= len(tokens) {
			return ""
		}
		v := tokens[*i]
		*i++

		l := evel(i, tokens, m)
		n := evel(i, tokens, m)

		if *i < len(tokens) && tokens[*i] == ")" {
			*i++
		}

		s := "(" + v + l + n + ")"
		m[s] = 1

		return s
	} else {
		k := tokens[*i]
		*i++
		return k
	}
}

func Polish(src string) int {
	m := make(map[string]int)

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Mode = scanner.ScanIdents

	var tokens []string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, s.TokenText())
	}

	idx := 0
	evel(&idx, tokens, m)
	return len(m)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var fullInput strings.Builder

	for sc.Scan() {
		fullInput.WriteString(sc.Text())
	}

	expr := strings.TrimSpace(fullInput.String())
	expr = strings.ReplaceAll(expr, " ", "")
	expr = strings.ReplaceAll(expr, "\t", "")
	expr = strings.ReplaceAll(expr, "\n", "")
	expr = strings.ReplaceAll(expr, "\r", "")

	if expr != "" {
		fmt.Println(Polish(expr))
	}
}
