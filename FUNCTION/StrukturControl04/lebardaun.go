package main
import "fmt"
func main(){
	var n, m, hasil int
	fmt.Scan(&n)
	m = n
	hasil = 0
	for i := 1; i <= m; i++ {
		fmt.Scan(&n)
		if n > hasil {
			hasil = n
		}
	}
	fmt.Println(hasil)
}
// package main
// import "fmt"
// func main(){
// 	var n, m, hasil int
// 	fmt.Scan(&n)
// 	m = n
// 	hasil = 0
// 	for i := 1; i <= m; i++ {
// 		fmt.Scan(&n)
// 		if n > hasil {
// 			hasil = n
// 		}
// 	}
// 	fmt.Println(hasil)
// }