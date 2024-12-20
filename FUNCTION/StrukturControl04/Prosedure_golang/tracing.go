package main

import "fmt"

func simpleCacl(b int, c *int){
	var a *int
	*a = 10 + b - *c
	*c = *c + 4
}
func main(){
	var x,y,z int
	fmt.Scan(&x,&y,&z)
	simpleCacl(y,&z)
	fmt.Print(y,z)

	fmt.Scan(&y)
	simpleCacl(z,&y)
	fmt.Print(x,y,z)
}