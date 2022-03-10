package main

import (
    "strings"
    "bufio"
    "fmt"
    "os"

    "github.com/Degen-dev/encode-decode/src"
)

var (
    scanner = bufio.NewScanner(os.Stdin)
)

func main() {

    fmt.Print("\033[H\033[2J") //clears terminal
    fmt.Print("Would you like to use base64 or base32: ")

    scanner.Scan()
    input := scanner.Text()

    if input == strings.ToLower("base64") {
	      src.Bs64()
    } else if input == strings.ToLower("base32") {
	      src.Bs32()
    } else {
	      fmt.Println("Dude, that wasn't an option.")
    }

}
