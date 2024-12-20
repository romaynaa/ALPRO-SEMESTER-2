package main
import "fmt"
func main(){
	var p,s,x,y,b,push,sit int
	fmt.Scan(&p,&s,&x,&y,&b)
	for i := 1; i < b * 30; i++ {
		if i % 2 != 0 && i % x != 0 && i % y != 0 {
			push += p
		}else if i % 2 == 0 && i % x != 0 && i % y != 0 {
			sit += s
		}
	}
	fmt.Println(push, sit)
}