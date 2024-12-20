package main
import "fmt"
func main(){
	 var n, ganjil, hasil int
	 fmt.Scan(&n)
	 hasil = 0
	 for i := 0; i < n; i++ {
		ganjil = 2 * i + 1
		hasil += ganjil
	 }
	 fmt.Println(hasil)
}