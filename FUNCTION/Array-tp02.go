package main

import "fmt"

const NMAX = 5

type tabInt struct {
	info[NMAX + 1] int
	n int
}

func main() {
	var nilaiAkhir tabInt
	// Melakukan pembacaan nilai akhir sebanyak 6x
	bacaNilai(&nilaiAkhir, 6)
	
	// Melakukan pencetakan nilai akhir
	cetakNilai(nilaiAkhir)
	
}

func bacaNilai(NA *tabInt, max int) {
	/*
		IS: NA.info[i] adalah field untuk menampung data nilai akhir, sedangkan 
		    NA.n untuk menampung banyaknya elemen data. Kedua field itu 
		    terdefinisi sembarang yang berarti bisa kosong atau berisi nilai.
		FS: Field nilai akhir (NA.info[i]) bertambah satu data selama belum melewati 
		    kapasitas maksimum array. Banyaknya elemen data (NA.n) bertambah satu
		    selama belum melewati kapasitas maksimum array.
	*/
	for i := 0; i < max; i++{
		fmt.Scan(&NA.info[NA.n])
		NA.n++
	}

}

func cetakNilai(NA tabInt) {
	/*
		IS: Nilai akhir (NA) terdefinisi sembarang, yang berarti bisa kosong atau 
		    berisi nilai.
		FS: Seluruh elemen tercetak di layar
	*/
	for i := 0; i < NA.n - 1; i++{
		fmt.Printf("%d ",NA.info[i])
	}
	fmt.Println()
}