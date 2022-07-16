package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BoolPrompt(label string, def bool) bool {
	choices := "Y/n"
	if !def {
		choices = "y/N"
	}

	r := bufio.NewReader(os.Stdin)
	var s string

	fmt.Fprintf(os.Stdout, "%s (%s) ", label, choices)
	s, _ = r.ReadString('\n')
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "y" || s == "yes" {
		return true
	} else if s == "n" || s == "no" {
		return false
	}

	return def
}
