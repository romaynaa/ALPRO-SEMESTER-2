package main

import "fmt"

func main(){
	var x int
	var status bool
	fmt.Scan(&x)
	status = true

	if x == 1 {
		status = false
	}else{
	for i := 2; i < (x-1); i++ {
		if x % i == 0 {
			status = false
		}
	}
}
	fmt.Println(status)
}
