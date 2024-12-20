package main
import "fmt"
func main(){
	var n, max, min, jumlah, i, end int
	var stop bool
	var rata float64
	i = 0
	stop = false
	for !stop {
		fmt.Scan(&n) 
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
		if n == 0 {
			end += 1
		}else if n != 0 {
			end = 0
		}
		jumlah += n
		i++
		stop = end == 2
	}
	rata = float64(jumlah) / float64(i-1)
	fmt.Println(max)
	fmt.Println(min)
	fmt.Println(rata)
}