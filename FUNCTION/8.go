package main 

import "fmt"

func gcd(a,b int)int{
	if b == 0{
		return a
	}
	return gcd(b, a%b)
}
func main(){
	var a,b int
	fmt.Scan(&a,&b)
	fmt.Println(gcd(a,b))
}