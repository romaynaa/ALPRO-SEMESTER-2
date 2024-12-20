package main 

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	printTriangleMaster(n)
}
/*func printTriangle(n int){
	if n == 1{
		fmt.Println("*")
	}else{
		printTriangle(n-1)
		for i := 0; i < n-1; i++{
			fmt.Print("*")
		}
		fmt.Println("*")
	}
}*/
func printTriangle(n int){
	for i := 0; i <= n; i++{
		for j := 0; j <= n; j++{
			fmt.Print("*")
			if i == j{
				fmt.Println()
			}
		}
	}
}
func printTriangleMaster(n int){
	var i int = 1
	printTriangleSlave(n,i)
}
func printTriangleSlave(n,i int){
	var j int
	if !(i < n){
		j = 1
		for j <= i{
			fmt.Print("*")
			if i == j{
				fmt.Println()
			}
			j++
		}
	}else{
		j = 1
		for j <= i{
			fmt.Print("*")
			if i == j{
				fmt.Println()
			}
			j++
		}
		i++
		printTriangleSlave(n,i)
	}
}