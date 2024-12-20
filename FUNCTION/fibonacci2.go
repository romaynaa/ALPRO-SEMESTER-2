package main 

import "fmt"

func main(){
	var n, f1, f2, fn int
	fmt.Scan(&n)
	f1 = 0
	f2 = 1
	fn = f1 + f2
	for i := 1; i <= n; i++ {
		fn = f1 + f2
		f2 = f1
		f1 = fn
	fmt.Println(fn)
	}
}
func fibonacci(n int)int{
	return hasil := hasil + n
}