package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func main() {

    fmt.Print("\033[H\033[2J") //clears terminal
	fmt.Print("Would you like to encode (e) or decode (d): ")
	scanner.Scan()
	input := scanner.Text()
    
	if input == "e" {
		encode()
	} else if input == "d" {
		decode()
	} else {
		fmt.Println("Dude, that wasn't an option.")
	}

}

func encode() {

	fmt.Print("Enter what you would like to be encoded: ")
    scanner.Scan()

	input := scanner.Text()
	Enc := base64.StdEncoding.EncodeToString([]byte(input))

	fmt.Println("Your encoded text is:", Enc)

}

func decode() {

	fmt.Print("Enter what you would like to be decoded: ")
    scanner.Scan()

	input := scanner.Text()
	Dec, _ := base64.StdEncoding.DecodeString(input)

	fmt.Println("Your decoded text is:", (string(Dec)))
	
}
