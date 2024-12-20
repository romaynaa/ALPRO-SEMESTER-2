package main 
import "fmt"
func main(){
	// var n, digit, hasil int
	// fmt.Scan(&n)\
	// for n > 0 {
	// 	if digit = digit / 10 
	// 		hasil = hasil < digit
	// 		n = n % 10 {
	// 			fmt.Println(asending)
	// 	}else if digit = digit / 10 
	// 		hasil = hasil > digit
	// 		n = n % 10 {
	// 		fmt.Println(disending)
	// 	}
	// }
	var n, ganjil, genap, tot_ganjil, tot_genap, bil int
	fmt.Scan(&n, &bil)
	ganjil = 0
	genap = 0
	for i := 1; i <= n; i++ {
		if bil % 2 == 0 && bil <= 15 {
			tot_genap = tot_genap + bil
			genap += 1
		}else if bil % 2 != 0 && bil <= 15 {
			tot_ganjil = tot_ganjil + bil
			ganjil += 1
		}
	}
	fmt.Println(ganjil, genap, tot_genap, tot_ganjil)
}