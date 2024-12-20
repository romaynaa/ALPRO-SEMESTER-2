package main 

import "fmt"

func fibonacci(n,f1,f2,fn int)int{
	f1 = 0
	f2 = 1
	fn = f1 + f2
	for i := 1; i <= n; i++ {
		fn = f1 + f2
		f2 = f1
		f1 = fn
	}
	return fn 	
}
func main(){
	var n, f1, f2, fn int
	fmt.Scan(&n)
	fmt.Println(fibonacci(n,f1,f2,fn))
}