package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	fmt.Println("dummy address:", randomdata.Address())
	fmt.Println("dummy currency:", randomdata.Currency())
	fmt.Println("dummy email:", randomdata.Email())
	fmt.Println("dummy phone number:", randomdata.PhoneNumber())
}
