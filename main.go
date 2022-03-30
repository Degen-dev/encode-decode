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

    switch string(strings.ToLower(scanner.Text())) {
    	case "base64":
			src.Bs64()
    	case "base32":
        	src.Bs32()
    	default:
        	fmt.Println("Dude, that wasn't an option.")
    }
}
