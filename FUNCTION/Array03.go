//PALINDROM
package main

import "fmt"

const NMAX int = 256
type tabInt[NMAX] int

func isiArray(info *tabInt, n *int){
	for i := 0; i < *n; i++{
		fmt.Scan(&info[i])
	}
}
func reverseArray(info *tabInt, n *int){
	
}