package main

import "fmt"

func Add(pp **int, x int) {
	**pp += x
}

func main() {
	var a int = 10    // Declare an integer variable
	var p *int = &a   // Pointer to the variable 'a'
	var pp **int = &p // Pointer to the pointer 'p'
	fmt.Printf("a with value of %v\n", a)
	fmt.Printf("&a with value of %v\n", &a)
	fmt.Printf("p with value of %v\n", p)
	fmt.Printf("*p with value of %v\n", *p)
	fmt.Printf("pp with value of %v\n", pp)
	fmt.Printf("*pp with value of %v\n", *pp)
	fmt.Printf("**pp with value of %v\n", **pp)
	fmt.Println("Now let's add 32 to pp")
	fmt.Println()
	fmt.Println(`func Add(pp **int, x int) {
	**pp += x
}
Add(pp, 32)`)
	fmt.Println()
	Add(pp, 32)

	fmt.Printf("a with value of %v\n", a)
	fmt.Printf("&a with value of %v\n", &a)
	fmt.Printf("p with value of %v\n", p)
	fmt.Printf("*p with value of %v\n", *p)
	fmt.Printf("pp with value of %v\n", pp)
	fmt.Printf("*pp with value of %v\n", *pp)
	fmt.Printf("**pp with value of %v\n", **pp)
}
