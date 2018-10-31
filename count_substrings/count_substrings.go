package main

import (
	"fmt"
	"strings"
	"time"
)

var week time.Duration

func main() {
	str := "Hello, how is it going, Hugo?"
	manyG := "gggggggggg"

	fmt.Printf("Number if H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is:", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))

	var arr [4]byte
	slice := []byte("abcdefg")
	copy(arr[:], slice)
	fmt.Println(arr)

}
