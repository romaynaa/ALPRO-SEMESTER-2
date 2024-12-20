package main
import "fmt"
func main(){
	var jabatan string
	var masa, anak, gaji int

	fmt.Scan(&jabatan, &masa, &anak)
	
	if jabatan=="staf" {
		if masa < 5 {
			gaji = 4000
}else if masa > 10 {
	gaji = 5000
	if anak >= 1 && anak <= 3 {
		gaji = 5000 + (anak * 100)
	}else if anak > 3 {
		gaji = 5000 + (0 * 100)
	}
}
	}
fmt.Println(gaji)
}