package src

import (
    "strings"
	"encoding/base64"
	"fmt"
)

func Bs64() {
  	fmt.Print("Would you like to encode or decode: ")

	scanner.Scan()
	input := strings.ToLower(scanner.Text())

	switch string(input) {
		case "encode":
		encode()
		case "decode":
		decode()
		default:
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
	Dec, err := base64.StdEncoding.DecodeString(input)

	fmt.Println("Your decoded text is:", (string(Dec)))

	if err != nil {
		fmt.Println("Dude, you need to input base64.")
	}
}
