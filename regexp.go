package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("ego"))
	fmt.Println(regex.MatchString("bayu"))
	fmt.Println(regex.MatchString("Eko"))

	fmt.Println(regex.FindAllString("eko edo egi ego e1o eto eKo", 10))
}