// package main

// import "fmt"

// func main(){
// 	var a,b int
// 	fmt.Scan(&a,&b)
// 	for i := a; ; i++{
// 		if i % a == 0 && i % b == 0{
// 			fmt.Println(i)
// 			break
// 		}
// 	}
// }
package main 

import "fmt"

func gcd(a,b int)int{
	if b == 0{
		return a
	}
	return gcd(b, a%b)
}
func lcm(a,b int)int{
	return a*b/gcd(a,b)
}
func main(){
	var a,b int
	fmt.Scan(&a,&b)
	fmt.Println(lcm(a,b))
}