package main

import "fmt"

const NMAX int = 10

type tabInt[NMAX] int

func main() {
	var data tabInt
	var nData int
	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println(jumlah(data,nData), rata_rata(data,nData))
}
func baca(A *tabInt, n *int) {
	/*
		IS: Array A dan n terdefinisi sembarang
		Proses: Setiap bilangan yang dibaca ditampung dalam sebuah variabel. 
		    Jika bilangan negatif, maka ubah menjadi bilangan positif dan
		    masukan ke dalam array. Pembacaan berakhir jika terbaca bilangan 0.
		FS: Array A sebanyak n elemen berisi bilangan positif
	*/
	var num int
	var i int = 0
	fmt.Scan(&num)
	for num != 0 && i < NMAX{
		A[i] = num
		i += 1
		fmt.Scan(&num)
	}
	*n = i
}
func cetak(A tabInt, n int) {
	/*
		IS: Array A terdefinisi sebanya kn elemen
		FS: Array A tercetak di layar
	 */
	for i := 0; i < n; i++{
		if A[i] > 0{
			fmt.Print(A[i]," ")
		}
	}
	fmt.Println()
}

func jumlah(A tabInt, n int) int {
	/* Mengembalikan jumlah dari nilai bilangan pada elemen array A */
	var jumlah int
	for i := 0; i < n; i++{
		if A[i] < 0{
			jumlah -= A[i]
		}else if A[i] > 0{
			jumlah += A[i]
		}
	}
	return jumlah
}

func rata_rata(A tabInt, n int) float64 {
	/* Mengembalikan rata-rata bilangan */
	return float64(jumlah(A,n))/float64(n)
}