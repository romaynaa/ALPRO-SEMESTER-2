package main
import "fmt"
func main(){
	var biner int
	var desimal, base, digit int
	desimal = 0
	base = 1
	fmt.Scan(&biner)
	for biner != 0 {
		digit = biner % 10 
		desimal = desimal + digit * base
		biner = biner / 10
		base = base * 2
	}
	fmt.Println(desimal)
}