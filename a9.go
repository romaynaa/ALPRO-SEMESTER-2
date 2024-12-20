package main
import "fmt"
func main(){
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if b < a {
		a,b = b,a
	}
	if c < a{
		a,c = c,a
	}
	if c < b {
		b,c = c,b
	}
	fmt.Println(a,b,c)
}