package main 

import "fmt"

func main () {
	var i = 2
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("%v %v %v\n",i,f,u)
}
