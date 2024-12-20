package main

import 
	"fmt"

/*
Tipe bentukan struktur tDrakor dengan atribut:
	judul (string), rating (real), jumlah episode (integer),
	dan durasi (integer)
*/
type tDrakor struct {
	judul string
	rating float64
	jumlah_episode, durasi int
}

// Konstanta NMAX dengan nilai 5
const NMAX int = 5

// Tipe alias tabDrakor untuk array of tDrakor dengan ukuran NMAX
type tabDrakor [NMAX]tDrakor

func main() {
	// Deklarasi variabel drakor bertipe array of tDrakor
	var drakor tabDrakor

	// Deklarasi variabel nDrakor bertipe integer untuk menampung banyaknya data drakor
	var nDrakor int

	// Pemanggilan prosedur bacaDataDrakor
	bacaDataDrakor(&drakor, &nDrakor)

	// Pemanggilan prosedur cetakDataDrakor.
	fmt.Println("Data sebelum diurutkan:")
	cetakDataDrakor(drakor, nDrakor)
	fmt.Println()

	// Pemanggilan prosedur urutMenurun
	urutMenurun(&drakor, nDrakor)

	// Pemanggilan prosedur cetakDataDrakor.
	fmt.Println("Data setelah diurutkan menurun berdasarkan rating:")
	cetakDataDrakor(drakor, nDrakor)
	fmt.Println()

	// Pemanggilan prosedur urutMenaik
	urutMenaik(&drakor, nDrakor)

	// Pemanggilan prosedur cetakDataDrakor.
	fmt.Println("Data setelah diurutkan menaik berdasarkan durasi:")
	cetakDataDrakor(drakor, nDrakor)
}

func bacaDataDrakor(A *tabDrakor, n *int) {
	/*
		IS: Array A sebanyak n terdefinisi sembarang
		Proses: 1) Membaca n. Nilai n tidak boleh melebihi NMAX.
				2) Membaca n data array A dengan atribut-atributnya.
		FS: Array A sebanyak n elemen berisi nilai
	*/
	fmt.Scan(n)
	if *n > NMAX{
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i].judul,&A[i].rating,&A[i].jumlah_episode,&A[i].durasi )
	}
}

func cetakDataDrakor(A tabDrakor, n int) {
	/*
		IS: Terdefinisi array A sebanyak n elemen
		FS: Tercetak di layar data array A dengan format
			" Judul    Rating       Jum Ep   Durasi
			 <judul> <rating> <jum episode> <durasi>
			  ...	  ...		...			 ...."
			  Gunakan format string dengan kolom 20, 6, 6, dan 6
			  untuk masing-masing judul, rating, jumlah episode,
			  dan durasi. 
	*/
	formatStr := "%20s %6s %6s %6s\n"
	fmt.Printf(formatStr, "Judul", "Rating", "Jum Ep", "Durasi")
	for i := 0; i < n; i++{
		fmt.Printf(formatStr, A[i].judul, fmt.Sprintf("%.1f", A[i].rating), fmt.Sprintf("%d",A[i].jumlah_episode), fmt.Sprintf("%d", A[i].durasi))
	}
}

func urutMenaik(A *tabDrakor, n int) {
	/*
		IS: Terdefinisi Array A sebanyak n elemen
		FS: Array A sebanyak n elemen terurut menaik berdasarkan
			durasi dengan menggunakan algoritma insertion sort
	*/
	for i := 0; i < n; i++{
		temp := A[i]
		j := i-1
		for j >= 0 && A[j].durasi > temp.durasi{
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
}

func urutMenurun(A *tabDrakor, n int) {
	/*
		IS: Terdefinisi Array A sebanyak n elemen
		FS: Array A sebanyak n elemen terurut menurun berdasarkan
			rating dengan menggunakan algoritma selection sort
	*/
	for i := 0; i < n; i++{
		maxIdx := i
		for j := i+1; j < n; j++{
			if A[j].rating > A[maxIdx].rating{
				maxIdx = j
			}
		}
		A[i], A[maxIdx]= A[maxIdx], A[i]
		
	}
}