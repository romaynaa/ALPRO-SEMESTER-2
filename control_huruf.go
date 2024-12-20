package main
import "fmt"
func main(){
	var ada bool
	var huruf rune
	// for n <= 10 {
	// 	if huruf == 'k' || huruf == 'q' {
	// 		ada = true
	// 	}else{
	// 		ada = false
	// 	}
	
	// }
	// fmt.Println(ada)
	// // n = 1
	// ada = false
	// for n <= 10 && !ada {
	// 	fmt.Scanf("%c", &huruf)
	// 	ada = huruf == 'k' || huruf == 'q'
	// 	n++
	// }
	// fmt.Println(ada)
	ada = false
	for n := 1; n <= 10 && !ada; n++ {
		fmt.Scanf("%c", &huruf)
		ada = huruf == 'k' || huruf == 'q'
	}
	fmt.Println(ada)
}
