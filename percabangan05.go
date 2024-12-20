package main

import "fmt"

func main(){
	var nilai int
	fmt.Scan(&nilai)

	if nilai >= 0 && nilai <= 100 {
		if nilai >= 80 {
			fmt.Println("A") 
		}else if nilai >= 70 && nilai < 80{
			fmt.Println("AB")
		}else if nilai >= 60 && nilai < 70{
			fmt.Println("B")
		}else if nilai >= 50 && nilai < 60{
			fmt.Println("C")
		}else if nilai >= 40 && nilai < 50{
			fmt.Println("D")
		}oooelse if nilai < 40{
			fmt.Println("E")
		}
	}else{
		fmt.Println("Data Invalid")
	}
}