package src

import (
  "os"
  "bufio"
	"encoding/base32"
	"fmt"
)

var (
    scanner = bufio.NewScanner(os.Stdin)
)

func Bs32() {

	fmt.Print("Would you like to encode (e) or decode (d): ")

	scanner.Scan()
	input := scanner.Text()

	if input == "e" {
		Bs32E()
	} else if input == "d" {
		Bs32D()
	} else {
		fmt.Println("Dude, that wasn't an option.")
	}

}

func Bs32E() {

	fmt.Print("Enter what you would like to be encoded: ")
	scanner.Scan()

	input := scanner.Text()
	bs32Enc := base32.StdEncoding.EncodeToString([]byte(input))

	fmt.Println("Your encoded text is:", bs32Enc)

}

func Bs32D() {

	fmt.Print("Enter what you would like to be decoded: ")
	scanner.Scan()

	input := scanner.Text()
	bs32Dec, err := base32.StdEncoding.DecodeString(input)

	fmt.Println("Your decoded text is:", string(bs32Dec))

	if err != nil {
		fmt.Println("Dude, you need to input base32.")
	}

}
