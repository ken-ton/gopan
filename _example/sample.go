package main

import (
	"fmt"
	"github.com/ken-ton/gopan"
)

func main() {
	//pan := "4111-1111-1111-1111"
	//pan := "3566 0020 2036 0505"
	//pan := Generate()
	pan := gopan.Generate("AMERICAN EXPRESS")

	hidden := gopan.GetHiddenPan(pan)
	valid := gopan.IsValid(pan)
	brand := gopan.GetBrand(pan)

	fmt.Printf("PAN: %s\n", pan)
	fmt.Printf("Hidden: %s\n", hidden)
	fmt.Printf("Length: %d\n", len(pan))
	fmt.Printf("Valid: %t\n", valid)
	fmt.Printf("Brand: %s\n", brand)
}
