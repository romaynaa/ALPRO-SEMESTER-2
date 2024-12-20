package main

import "fmt"

const NMAX int = 11

type pemain struct {
	nama, nomorPunggung, posisi string
	tinggi                      int
}

// tabPemain adalah alias array pemain dengan maks elemen NMAX
type tabPemain [NMAX] pemain

func main() {
	var timnas tabPemain
	var nPemain int
	nPemain = 0
	// baca data
	bacaData(&timnas, &nPemain)
	// cetak data pemain
	cetakPemain(timnas, nPemain)
	// cetak nama pemain tertinggi
	cetakNamaPemainTertinggi(timnas, nPemain)
	// cetak nama pemain terendah
	cetakNamaPemainTerendah(timnas, nPemain)
}

func bacaData(A *tabPemain, n *int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi sembarang
		Proses: Membaca nama, nomor punggung, posisi, dan tinggi badan
			 	Jika array belum penuh dan nama bukan "none", maka nama, nomor punggung, posisi,
				dan tinggi badan dimasukkan ke dalam array.
		FS: Array A sebanyak n elemen berisi nilai
	*/
	var nm, no, psi string
	var t int
	var i int = 0
	fmt.Scanln(&nm,&no,&psi,&t)
	for nm != "none" && i < NMAX{
		A[i].nama = nm
		A[i].nomorPunggung = no
		A[i].posisi = psi
		A[i].tinggi = t
		i++
		fmt.Scanln(&nm,&no,&psi,&t)
	}
	*n = i
	if *n > NMAX{
		*n = NMAX
	}
}

func cetakPemain(A tabPemain, n int) {
	/*
		IS: Array A dengan banyak elemen n terdefinisi
		FS: Tercetak di layar elemen array A sebanyak n dengan format:
			"Data pemain:
			<nama1> <nomorPunggung1> <posisi1> <tinggi1>
			<nama2> <nomorPunggung2> <posisi2> <tinggi2>
			...
			<naman> <nomorPunggungn> <posisin> <tinggin>"
	*/
	fmt.Println("Data Pemain :")
	for i := 0; i < n; i++{
		fmt.Printf("%s %s %s %d\n", A[i].nama, A[i].nomorPunggung, A[i].posisi, A[i].tinggi)
	}
}

func cetakNamaPemainTertinggi(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan tertinggi dengan format:
	       "Pemain dengan badan tertingggi: <nama>"
		   Asumsi: Hanya ada 1 pemain dengan badan tertinggi
	*/
	var nm string
	var tertingggi int
	tertingggi = A[0].tinggi
	for i := 0; i < n; i++{
		if A[i].tinggi > tertingggi{
			nm = A[i].nama
			tertingggi = A[i].tinggi
		}
	}
	fmt.Printf("Pemanin dengan badan tertinggi: %s\n", nm)
}

func cetakNamaPemainTerendah(A tabPemain, n int) {
	/* IS: Array A sebanyak n elemen terdefinisi
	   FS: Mencetak nama pemain dengan badan terendah dengan format:
	       "Pemain dengan badan terendah: <nama>""
		   Asumsi: Hanya ada 1 pemain dengan badan terendah
	*/
	var nm string
	var terendah int
	for i := 0; i < n; i++{
		terendah = A[i].tinggi
		if A[0].tinggi < terendah{
			nm = A[0].nama
			terendah = A[i].tinggi
		}
	}
	fmt.Printf("Pemanin dengan badan terendah: %s\n", nm)
}
