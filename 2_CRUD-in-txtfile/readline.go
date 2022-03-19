package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("simple.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	myScanner := bufio.NewScanner(f)
	myScanner.Split(bufio.ScanLines)

	var fileText []string

	for myScanner.Scan(){
		fileText = append(fileText, myScanner.Text())
	}

	f.Close()

	for _, i := range fileText {
		fmt.Println(i)
	}

	fmt.Println(fileText)
}
