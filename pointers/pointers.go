package main

import "fmt"

func main() {
	var age int = 32

	var agePointer *int = &age

	fmt.Println("original pointer address: ", &age)
	fmt.Println("is adult?", isAdultInYears(age))                                                                        // creates a copy, we send the value
	fmt.Println("is adult (use common function, but send the value from the age pointer)?", isAdultInYears(*agePointer)) // makes a copy even though we send a pointer, 
	// it's because we need to create a value within the function scope
																	
	fmt.Println("is adult (using separate function which takes a age-pointer)?", isAdultInYearsPointer(agePointer)) // sends a pointer, does not create a separate value within the function scope
	fmt.Println("is adult (using separate function which takes age as pointer &age)?", isAdultInYearsPointer(&age)) // sends pointer as well, does not create a separate value

	// original pointer address:  0x14000096008
	// age value sent:  32 pointer:  0x14000096010
	// is adult? true
	// age value sent:  32 pointer:  0x14000096018
	// is adult (use common function, but send the value from the age pointer)? true
	// age by pointer is:  32 pointer:  0x14000096008
	// is adult (using separate function which takes a age-pointer)? true
	// age by pointer is:  32 pointer:  0x14000096008
	// is adult (using separate function which takes age as pointer &age)? true

}

// we create a copy of age variable and store it in the memory twice
func isAdultInYears(age int) bool {
	fmt.Println("age value sent: ", age, "pointer: ", &age)
	return age-18 > 0
}

// we do not create a separate copy of age, instead we use the pointer which points the same place in memory
// could be useful to prevent unnecessary copying the values
func isAdultInYearsPointer(age *int) bool {
	fmt.Println("age by pointer is: ", *age, "pointer: ", age)
	return *age-18 > 0
}
