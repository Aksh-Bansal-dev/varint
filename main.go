package main

import "fmt"

func main() {
	num := int64(12020)
    fmt.Println("Number:",num)
	encodedNum := EncodeInt64(num)
    fmt.Println("Encoded:",encodedNum)
	decodedNum := DecodeInt64(encodedNum)
    fmt.Println("Decoded:",decodedNum)
}
