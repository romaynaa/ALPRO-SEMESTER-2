package main
import "fmt"
func main(){
	var n, positif, negatif int
	fmt.Scan(&n)

	for n != 0 {
		if n > 0 {
			positif += 1
		}else if n < 0 {
			negatif += 1
		}
		fmt.Scan(&n)
	}
	fmt.Println(positif,negatif)
}