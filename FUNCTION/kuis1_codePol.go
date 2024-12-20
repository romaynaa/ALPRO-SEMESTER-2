// package main

// import "fmt"
// func main(){
// 	var a,b,c int
// 	a = 12
// 	b = a + 15
// 	c = b - (2*a)
// 	a = c
// 	fmt.Println(a)

// }
package main

import "fmt"

func main(){
	var temp string
	dataSiswa = array[4]  string
	var n int = dataSiswa
	dataSiswa[0] = "Januari"
	dataSiswa[0] = "Febri"
	dataSiswa[0] = "Agus"
	dataSiswa[0] = "Novi"
	dataSiswa[0] = "Desi"

	for i := 0; i < n; i++{
		for k := 0; k < i; k++{
			if posisiAbjad(dataSiswa[k]) > posisiAbjad(dataSiswa[k+1]){
				temp = dataSiswa[k]
				dataSiswa[k] = dataSiswa[k+1]
				dataSiswa[k+1] = temp
			}
		}
	}

}