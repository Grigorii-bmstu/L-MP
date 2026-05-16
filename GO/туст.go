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
	// Если мы зашли в скобки
	if tokens[*i] == "(" {
		*i++                 // Пропускаем "("
		v := tokens[*i]      // Запоминаем операцию (+, -, *)
		*i++                 // Пропускаем операцию, чтобы палец стоял на первом числе
		l := evel(i, tokens) // Вычисляем первый аргумент
		n := evel(i, tokens) // Вычисляем второй аргумент
		*i++                 // Пропускаем закрывающую скобку ")"
		if v == "+" {
			return l + n
		} else if v == "-" {
			return l - n
		} else if v == "*" {
			return l * n
		}
	} else {
		k, _ := strconv.Atoi(tokens[*i])
		*i++ // ОБЯЗАТЕЛЬНО двигаем палец дальше после чтения числа
		return k
	}
	return 0
}

func Polish(src string) int {
	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	var tokens []string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, s.TokenText())
	}
	if len(tokens) == 0 {
		return 0
	}
	idx := 0
	return evel(&idx, tokens)
}

func main() {
	// Читаем ввод из консоли
	reader := bufio.NewReader(os.Stdin)
	expr, _ := reader.ReadString('\n')
	expr = strings.TrimSpace(expr)

	if expr != "" {
		fmt.Println(Polish(expr))
	}
}
