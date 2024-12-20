package main

import "fmt"

func main(){
	var a, b, c, d, gol_banyak, gol_dikit int
	fmt.Scan(&a,&b,&c,&d)

	if a >= b && a >= c && a >= d {
		gol_banyak = a
	}else if b >= a && b >= c && b >= d {
		gol_banyak = b
	}else if c >= a && c >= b && c >= d {
		gol_banyak = c
	}else if d >= a && d >= b && d >= c {
		gol_banyak = d
	}
	
	if a <= b && a <= c && a <= d {
		gol_dikit = a
	}else if b <= a && b <= c && b <= d {
		gol_dikit = b
	}else if c <= a && c <= b && c <= d {
		gol_dikit = c
	}else if d <= a && d <= b && d <= c {
		gol_dikit = d
	}
	fmt.Println(gol_banyak, gol_dikit)
}