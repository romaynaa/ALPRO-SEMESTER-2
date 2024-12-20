package main
import "fmt"
func main(){
	var n, genap, hasil int
	fmt.Scan(&n)
	hasil = 0
	for i := 0; i < n; i++ {
		genap = 2 * i + 2
		hasil += genap
	}
	fmt.Println(hasil)
}