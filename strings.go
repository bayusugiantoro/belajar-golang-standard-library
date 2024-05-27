package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Bayu Sugiantoro", "Bayu"))
	fmt.Println(strings.Split("Bayu Sugiantoro", " "))
	fmt.Println(strings.ToLower("BAYU SUGIANTORO"))
	fmt.Println(strings.ToUpper("bayu sugiantoro"))
	fmt.Println(strings.Trim("      Bayu Sugiantoro      ", " "))
	fmt.Println(strings.ReplaceAll("Eko Kurniawan Bayu Sugiantoro", "Bayu", "Bagus"))
}