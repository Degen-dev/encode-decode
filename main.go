package main

import (
    "bufio"
    "fmt"
    "os"
)

var (
    scanner = bufio.NewScanner(os.Stdin)
)

func main() {

    fmt.Print("\033[H\033[2J") //clears terminal
    fmt.Print("Would you like to use base64 (64) or base32 (32): ")

    scanner.Scan()
    input := scanner.Text()

    if input == "64" {
	Bs64()
    } else if input == "32" {
	Bs32()
    } else {
	fmt.Println("Dude, that wasn't an option.")
    }

}
