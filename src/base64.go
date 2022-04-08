/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package src

import (
    "strings"
	"encoding/base64"
	"fmt"
)

func Bs64() {
  	fmt.Print("Would you like to encode or decode: ")
	scanner.Scan()

	switch string(strings.ToLower(scanner.Text())) {
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

	Enc := base64.StdEncoding.EncodeToString([]byte(scanner.Text()))

	fmt.Println("Your encoded text is:", Enc)
}

func decode() {
	fmt.Print("Enter what you would like to be decoded: ")
	scanner.Scan()

	Dec, err := base64.StdEncoding.DecodeString(scanner.Text())

	fmt.Println("Your decoded text is:", (string(Dec)))

	if err != nil {
		fmt.Println("Dude, you need to input base64.")
	}
}
