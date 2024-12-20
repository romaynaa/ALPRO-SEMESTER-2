package main

import "fmt"

func menu(){
	fmt.Println("-----------------------")
	fmt.Println("        M E N U        ")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. pengurangan")
	fmt.Println("3. perkalian")
	fmt.Println("4. pembagian")
	fmt.Println("5. Exit")
	fmt.Println("-----------------------")
}
func penjumlahan(a,b float64, hasil *float64){
	fmt.Scan(&a,&b)
	fmt.Println("Hasil penjumlahan: ", a+b)
}
func pengurangan(a,b float64, hasil *float64){
	fmt.Scan(&a,&b)
	fmt.Println("Hasil pengurangan: ", a-b)
}
func perkalian(a,b float64, hasil *float64){
	fmt.Scan(&a, &b)
	fmt.Println("Hasil perkalian: ", a*b)
}
func pembagian(a,b float64, hasil *float64){
	fmt.Scan(&a,&b)
	fmt.Println("Hasil pembagian: ", a/b)
}
func main(){
	var pilih int
	var x,y,hasil float64
	
	for {
		menu()
		fmt.Println("Pilih (1/2/3/4/5)? ")
		fmt.Scan(&pilih)
		switch pilih{
		case 1:
			penjumlahan(x,y,&hasil)
		case 2:
			pengurangan(x,y,&hasil)
		case 3:
			perkalian(x,y,&hasil)
		case 4:
			pembagian(x,y,&hasil)
		case 5:
			return
		}
	}
	fmt.Println(hasil)
}