package main

import (
	"bufio"
	"exercism/hamming"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter DNA strand 1: ")
	strand1, _ := reader.ReadString('\n')
	fmt.Print("Enter DNA strand 2: ")
	strand2, _ := reader.ReadString('\n')
	distance, err := hamming.Distance(strand1, strand2)
	if err != nil {
		fmt.Println(err)
	}
	if err == nil {
		fmt.Print("The hamming distance is: ")
		fmt.Println(distance)
	}
}
