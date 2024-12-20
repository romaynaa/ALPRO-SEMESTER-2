package main

import "fmt"

// Deklarsi konstanta 20 elemen
const NMAX int = 20

// Deklarasi tabInt dengan total NMAX elemen
type tabInt [NMAX]int

func main() {
    // Deklarasi variabel bertipe tabInt
	var Info tabInt
	// Deklarasi banyak elemen array
	var ndata int
	// Pembacaan data bilangan
	baca(&Info, &ndata)
	// Cetak elemen array
	cetakElemen(Info, ndata)
	// Cetak Info
	cetakInfo(Info, ndata)
}

func baca(A *tabInt, n *int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi sembarang
		Proses: Membaca bilangan bulat dan memasukkan bilangan bulat positif 
		        ke dalam array A.Pembacaan dilakukan selama terbaca 
		        bilangan positif dan n < NMAX.
		FS: Array A sebanyak n elemen berisi nilai
	*/
	var i int = 0
	var num int
	for i < NMAX{
		fmt.Scan(&num)
		if num <= 0{
			break
		}
		(*A)[i] = num
		i++
	}
	*n = i
	
}

func cetakElemen(A tabInt, n int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi
		FS: Tercetak di layar elemen array A sebanyak n dengan format:
			"Elemen array: <e1> <e2> <e3> ... <en>"
	*/
	fmt.Print("Elemen Array : ",)
	for i := 0; i < n; i++{
		fmt.Printf("%d ",A[i])
		fmt.Print(" ")
	}
	fmt.Println()
}

func maksimum(A tabInt, n int) int {
	/* Mengembalikan nilai elemen maksimum dari array A dengan banyak elemen n */
	var max int = A[0]
	for i := 0; i < n; i++{
		if A[i] > max{
			max = A[i]
		}
	}
	return max
}

func minimum(A tabInt, n int) int {
	/* Mengembalikan nilai elemen minimum dari array A dengan banyak elemen n */
	var min int = A[0]
	for i := 0; i < n; i++{
		if A[i] < min{
			min = A[i]
		}
	}
	return min
}

func cetakInfo(A tabInt, n int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi
		FS: Nilai maksimum, nilai minimum, dan banyaknya elemen tercetak dengan format:
			"Nilai maksimum: <max_value>
			 Nilai minimum: <min_value>
			 Banyak elemen: <n>"
	*/
	fmt.Printf("Nilai maksimum : %d\n", maksimum(A,n))
	fmt.Printf("Nilai minimum : %d\n", minimum(A,n))
	fmt.Printf("Banyak Elemen : %d\n", n)
}
