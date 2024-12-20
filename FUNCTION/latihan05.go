package main 

import "fmt"

func main(){
	var a,b,c int
	fmt.Scan(&a,&b,&c)
	if bentukSegitiga(a,b,c) == true{
		fmt.Println("Segitiga")
	}else{
		fmt.Println("Bukan segitiga")
	}

}
func bentukSegitiga(x,y,z int)bool{
	var hasil bool
	hasil =  x + y > z && x + z > y && y + z > x
	return hasil
}
