package main

import "fmt"

func main(){
	var jabatan string
	var masa, anak, gaji int
	fmt.Scan(&jabatan, &masa, &anak)

	if jabatan == "staf" && masa < 5 {
		gaji = 4000
	}else if jabatan == "staf" && masa > 10 && anak >= 1 && anak <= 3 {
		gaji = 5000 * (anak * 100)
	}else if jabatan == "staf" && masa > 10 && anak >= 3 {
		gaji = 5000 * (3 * 100)
	}else if jabatan == "staf" && masa >= 5 && masa <= 10 && anak >= 1 && anak <= 3 {
		gaji = 4000 + (anak * 100)
	}else if jabatan == "staf" && masa >= 5 && masa <= 10 && anak >= 3 {
		gaji = 4000 + (3 * 100)
	}
	if jabatan == "menejer" && masa > 10 && anak >= 1 && anak <= 3 {
		gaji = 10000 + (anak * 300)
	}else if jabatan == "menejer" && masa > 10 && anak >= 3 {
		gaji = 10000 + (3 * 300)
	}else if jabatan == "menejer" && masa <= 10 && anak >= 1 && anak <= 3 {
		gaji = 8500 + (anak * 300)
	}else if jabatan == "menejer" && masa <= 10 && anak >= 3 {
		gaji = 8500 + (3 * 300)
	}
	if jabatan == "direktur" && anak >= 1 && anak <= 3 {
		gaji = 20000 + (anak * 500)
	}else if jabatan == "direktur" && anak >= 3 {
		gaji = 20000 + (3 * 500)
	}
	fmt.Println(gaji)
}