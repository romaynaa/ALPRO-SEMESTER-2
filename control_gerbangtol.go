package main
import "fmt"
func main(){
	var a, b, c int
	var n string
	var stop bool
	stop = true	
	for stop {
		fmt.Scan(&n)
		if n == "A"{
			a += 1
		}else if n == "B" {
			b += 1
		}else if n == "C" {
			c += 1
		}else{
			stop = false
		}

	}
fmt.Println("Tipe A", a)
fmt.Println("Tipe B", b)
fmt.Println("Tipe C", c)

}