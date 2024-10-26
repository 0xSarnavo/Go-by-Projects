package main

import "fmt"

const s string = "constant"

func main() {

	fmt.Println("Printing s:",s)

	var a = "This"
	var b = "is a string"
	f := "This is also a string"
	fmt.Println("Printing a and b:",a,b)
	fmt.Println("Printing f:",f)

	var c,d = 1,2
	fmt.Println(c+c,d)
	fmt.Println(c==d)
	fmt.Println(c!=d)

	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string
	fmt.Printf("%v %f %v %q\n",smsSendingLimit,costPerSMS,hasPermission,username)
}
