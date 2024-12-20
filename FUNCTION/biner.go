package main 

import "fmt"

func main(){
	var biner, desimal, weight, digit int
	fmt.Scan(&biner)
	desimal = 0
	weight = 1
	for biner != 0 {
		digit = biner % 10
		desimal = desimal + digit * weight
		biner = biner / 10
		weight = weight * 2
		
	}
	fmt.Println(desimal)
}