package stringutil

import (
	"strings"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func CompareThese(a, b string) string {
	var c string
	if len(a) > len(b) {
		c = a + " has more characters than " + b + ".\n"
	}
	if len(a) < len(b) {
		c = a + " has fewer characters than " + b + ".\n"
	}
	if len(a) == len(b) {
		c = a + " and " + b + " are equal in length of characters.\n"
	}
	return c
}

func ConcatThese(a, b string) string {
	concatenated := a + b
	return concatenated
}
