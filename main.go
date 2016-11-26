package main

import "fmt"

func main() {
	names := [4]string{
		"Josh", "Pedro", "Leo", "Carol",
	}

	fmt.Println("Names array:", names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println("a slice:", a)
	fmt.Println("b slice:", b)
	fmt.Println("Changed item [0]  of slice b to Nathalia")
	b[0] = "Nathalia"
	fmt.Println("a slice:", a)
	fmt.Println("b slice:", b)
	fmt.Println("names array:", names)
}
