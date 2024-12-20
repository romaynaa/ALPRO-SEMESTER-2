package main

import "fmt"

func faktorial(n int)int{
	if n == 1 || n == 0{
		return 1
	}else{
		return n * faktorial(n-1)
	}
}
func main(){
	var a,b,c,d int
	var permutasi1, kombinasi1, permutasi2, kombinasi2 int
	fmt.Scan(&a,&b,&c,&d)
	if a >= c {
		permutasi1 = faktorial(a)/faktorial(a-c)
		kombinasi1 = faktorial(a)/(faktorial(c)*faktorial(a-c))
	}
	if  b >= d {
		permutasi2 = faktorial(b)/faktorial(b-d)
		kombinasi2 = faktorial(b)/(faktorial(d)*faktorial(b-d))
	}
	fmt.Println(permutasi1,kombinasi1,permutasi2,kombinasi2)
}