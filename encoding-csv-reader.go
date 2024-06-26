package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "eko,kurniawan,khannedy\n" +
		"budi,pratama,nugraha\n" +
		"joko,morro,diah\n" +
		"bayu,sugiantoro"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}