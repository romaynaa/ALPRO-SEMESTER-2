package main

import "fmt"

func nilaiSukuKe(n int)int{
	if n <= 3 {
		return 1
	}
	return nilaiSukuKe(n-1) + nilaiSukuKe(n-2) + nilaiSukuKe(n-3)
}
func main(){
	var n int
	fmt.Scan(&n)
	fmt.Println(nilaiSukuKe(n))
}