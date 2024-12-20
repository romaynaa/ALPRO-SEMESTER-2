package main

import "fmt"

// Deklarasi konstanta NMAX dengan nilai 10
const NMAX int = 10

// Deklarasi tipe alias tabInt untuk array bilangan bulat dengan 
// kapasitas maksimum NMAX
type tabInt [NMAX]int

func main() {

	// Deklarasi variabel
	var data tabInt
	var nData int
	var x2 int
	var idx int

	// Baca bilangan yang dicari
	fmt.Scan(&x2)

	// Baca data array
	bacaData(&data, &nData)

	// Cetak data
	cetakData(data, nData)

	// Pencarian bilangan tertentu dengan binary search
	fmt.Print("Hasil pencarian: ")
	binarySearch(data,nData,x2,&idx)

    // Jika idx tidak bernilai -1, maka bilangan yang dicari ditemukan  
	if idx != -1 {
		fmt.Printf("Bilangan %d ditemukan pada posisi ke-%d\n", x2, idx)
	} else {
		fmt.Println("Bilangan tidak ditemukan.")
	}
}

func bacaData(A *tabInt, n *int) {
	/*
		IS: Array bilangan bulat A dan n terdefinisi sembarang
		FS: Array bilangan bulat A terisi data sebanyak n elemen.
			Elemen array terisi bilangan bulat menaik (ascending). 
			Jika bilangan yang dibaca lebih kecil dari
			bilangan sebelumnya, data tidak masuk ke dalam array dan 
			pembacanaan terhenti.
	*/
	var bil int
	fmt.Scan(&bil)
	for *n < NMAX && bil != 0{
		A[*n] = bil
		*n = *n + 1
		fmt.Scan(&bil)
		if bil < A[*n-1] {
			bil = 0
		}
	}
}

func cetakData(A tabInt, n int) {
	/*
		IS: Array bilangan A dan banyak elemen n terdefinisi
		FS: Tercetak di layar bilangan dengan format:
			"Data bilangan: <e1> <e2> <e3> ..<en>"
	*/
	fmt.Printf("Data Bilangan:")
	for i := 0; i < n; i ++{
		fmt.Printf(" %d", A[i])
	}
	fmt.Printf("\n")
}

func binarySearch(A tabInt, n int, x int, idx *int) {
	/*
		IS: Array bilangan A sebanyak n elemen dan bilangan x yang 
		    dicari terdefinisi
		FS: idx berisi indeks lokasi x berada jika x ditemukan. 
		    Jika tidak ditemukan bernilai -1.
	*/
	var left, right, mid int
	left = 0
	right = n - 1
	mid = (left + right) / 2
	*idx = -1
	for left <= right && A[mid] != x{
		if x < A[mid] {
			right = mid - 1
		}else{
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	if A[mid] == x{
		*idx = mid + 1
	}
}