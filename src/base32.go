package src

import (
    "strings"
    "os"
    "bufio"
	"encoding/base32"
	"fmt"
)

var (
    scanner = bufio.NewScanner(os.Stdin)
)

func Bs32() {
	fmt.Print("Would you like to encode or decode: ")
	scanner.Scan()

	switch string(strings.ToLower(scanner.Text())) {
		case "encode":
			Bs32E()
		case "decode":
			Bs32D()
		default:
			fmt.Println("Dude, that wasn't an option.")
	}
}

func Bs32E() {
	fmt.Print("Enter what you would like to be encoded: ")
	scanner.Scan()

	bs32Enc := base32.StdEncoding.EncodeToString([]byte(scanner.Text()))

	fmt.Println("Your encoded text is:", bs32Enc)
}

func Bs32D() {
	fmt.Print("Enter what you would like to be decoded: ")
	scanner.Scan()

	bs32Dec, err := base32.StdEncoding.DecodeString(scanner.Text())

	fmt.Println("Your decoded text is:", string(bs32Dec))

	if err != nil {
		fmt.Println("Dude, you need to input base32.")
	}
}
