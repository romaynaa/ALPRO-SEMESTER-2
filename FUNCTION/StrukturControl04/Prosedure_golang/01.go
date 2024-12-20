// 
package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	cetak_rerata(N, 0, 0, 1)
}

func cetak_rerata(N, jumlah int, rata float64, i int) {
	if i > N {
		return
	}
	jumlah += i
	rata = float64(jumlah) / float64(i)
	fmt.Println(jumlah, rata)
	cetak_rerata(N, jumlah, rata, i+1)
}