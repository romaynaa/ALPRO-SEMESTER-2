package main 

import "fmt"

func main(){
	var s1, s2, s3, s4, s5 string
	fmt.Scan(&s1,&s2,&s3,&s4,&s5)

	if s1 == "kalah" && s2 == "kalah" && s3 == "kalah" && s4 == "kalah" && s5 == "kalah" {
		fmt.Println("dipecat")
	}else{
		fmt.Println("tidak dipecat")
	}
}