package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//  bufio:-Package bufio implements buffered I/O.
// It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that
// also implements the interface but provides buffering and some help for textual I/O.

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 to 5 point.")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for rating, ", input)

	// numRating =input +1 gives error because input is string
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your Rating: ", numRating+1)
	}

}
