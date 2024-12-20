package main 
import "fmt"

func main(){
	var a1,a2,a3,a4,maximum, minimum int
	fmt.Scan(&a1,&a2,&a3,&a4)
	mencariNilaiExtrem(a1,a2,a3,a4,&maximum,&minimum)
	fmt.Print(maximum,minimum)
}
func mencariNilaiExtrem(n1,n2,n3,n4 int, max, min *int){
	*max = n1 
	if n2 > *max{
		*max = n2
	}
	if n3 > *max{
		*max = n3
	}
	if n4 > *max{
		*max = n4
	}

	*min = n1
	if n2 < *min{
		*min = n2
	}
	if n3 < *min{
		*min = n3
	}
	if n4 < *min{
		*min = n4
	}

}
