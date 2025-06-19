package main

import "fmt"

func main() {
	var(
		i int
		f float64
		b bool
		s string 
	)
	fmt.Println("%v %v %v %q\n",i,f,b,s)
	fmt.Println()
	fmt.Printf("%v %v %v %q\n",i,f,b,s)
}

