/*package main
//TIPE 1
const NMAX int = 256
type tabInt [NMAX] int

func main(){
	var data tabInt
	var ndata int

	data[0] = 8
	ndata = 1
	fmt.Println(data[1], ndata)
}*/
//TIPE 2
/*package main
import "fmt"

const NMAX int = 1024
type tabInt struct{
	info [NMAX] int
	ndata int
}
func main(){
	var data tabInt
	data.info[0] = 8
	data.ndata = 1
	fmt.Println(data.info[0], data.ndata)
}*/

/*package main

import "fmt"

const NMAX int = 100
type tabInt[NMAX]byte

func main(){
	var ndata int
	var data tabInt

	readArray(&ndata, &data)
	cetakArray(ndata, data)
}
func readArray(ndata *int, data *tabInt){
	var kata byte
	var i int = 0
	for kata != '.' && i < NMAX{
		data[i] = kata
		i += 1
		fmt.Scanf("%c", &kata)
	}
	*ndata = i
}
func cetakArray(ndata int, data tabInt){
	for i := ndata; i > 0; i--{
		fmt.Printf("%c", data[i])
	}
}
*/
/*package main 

import "fmt"

const NMAX int = 256
type tabInt struct{
	info[NMAX] int
	ndata int
}
func main(){
	var d tabInt
	bacadata(&d)
	fmt.Printf("Jumlah bilangan = %d \n", hitungJumlah(d))
	fmt.Printf("Rata-rata bilangan = %.2f", rerata(d))

}
func bacadata(d *tabInt){
	fmt.Scan(&d.ndata)
	for i := 0; i < d.ndata; i++{
		fmt.Scan(&d.info[i])
	}
}
func hitungJumlah(d tabInt)int{
	var jumlah int = 0
	for i := 0; i < d.ndata; i++{
		jumlah += d.info[i]
	}
	return jumlah
}
func rerata(d tabInt)float64{
	return float64(hitungJumlah(d)) / float64(d.ndata)
}*/

/*package main

import "fmt"

const NMAX int = 256
type tabInt struct{
	info[NMAX] int
	ndata int
}
func main(){
	var d tabInt
	bacadata(&d)
	fmt.Printf("jumlah bilangan = %d\n", hitungJumlah(d))
	fmt.Printf("jumlah rerata = %.2f", hitungRerata(d))
}
func bacadata(d *tabInt){
	fmt.Scan(&d.ndata)
	for i := 0; i < d.ndata; i++{
		fmt.Scan(&d.info[i])
	}
}
func hitungJumlah(d tabInt)int{
	var jumlah int = 0
	for i := 0; i < d.ndata; i++{
		jumlah += d.info[i]
	}
	return jumlah
}
func hitungRerata(d tabInt)float64{
	var rerata float64
	rerata = float64(hitungJumlah(d))/float64(d.ndata)
	return rerata
}*/

/*package main

import "fmt"

const NMAX int = 10
type tabInt[NMAX] byte

func main(){
	var info tabInt
	var ndata int
	bacaArray(&info, &ndata)
	cetakArray(info, ndata)
}
func bacaArray(info *tabInt, ndata *int){
	var kata byte
	var i  int = 0
	for kata != '.' && i < NMAX{
		info[i] = kata
		i += 1
		fmt.Scanf("%c", &kata)
	}
	*ndata = i
}
func cetakArray(info tabInt, ndata int){
	/*for i := ndata; i >= 0; i--{
		fmt.Printf("%c", info[i])
	}*/
	/*for ndata > 0{
		fmt.Printf("%c", info[ndata-1])
		ndata = ndata - 1
	}
}
*/
/*
package main

import "fmt"

const NMAX int = 100
type tabInt[NMAX] byte
func main(){
	var info tabInt
	var ndata int
	bacaArray(&ndata, &info)
	cetakArray(ndata, info)

}
func bacaArray(ndata *int, info *tabInt){
	var kata byte
	var i int = 0
	for kata != '.' && i < NMAX{
		info[i] = kata
		i += 1
		fmt.Scanf("%c", &kata)
	}
	*ndata = i
}
func cetakArray(ndata int, info tabInt){
	/*for ndata >= 0{
		fmt.Printf("%c", info[ndata-1])
		ndata = ndata - 1
	}*/
	/*for i := ndata; i >= 0; i--{
		fmt.Printf("%c", info[i])
	}
}*/
/*package main

import "fmt"

const NMAX int = 250
type tabInt struct{
	info[NMAX] int
	ndata int
}
func main(){
	var d tabInt
	bacaArray(&d)
	fmt.Printf("jumlah bilangan = %d\n", hitungJumlah(d))
	fmt.Printf("jumlah rerata = %.2f\n", hitungRerata(d))

}
func bacaArray(d *tabInt){
	fmt.Scan(&d.ndata)
	for i := 0; i < d.ndata; i++{
		fmt.Scan(&d.info[i])
	}
}
func hitungJumlah(d tabInt)int{
	var jumlah int = 0
	for i := 0; i < d.ndata; i++{
		jumlah += d.info[i]
	}
	return jumlah
}
func hitungRerata(d tabInt)float64{
	return float64(hitungJumlah(d))/float64(d.ndata)
}*/

package main

import "fmt"
const NMAX int = 100
type tabInt [NMAX] penerbangan

type penerbangan struct{
	maskapai string
	bandara_asal string
	bandara_tujuan string
	hari string
	waktu_berangkat waktu
	waktu_tiba waktu
	durasi int
}
type waktu struct{
	jam int
	menit int
}
func main(){
	var info tabInt
	var ndata int
	fmt.Scan(&ndata)
	bacaArray(&info, ndata)
	lamaPerjalanan(&info, ndata)
	cetakArray(&info, ndata)
}
func bacaArray(info *tabInt, ndata int){
	fmt.Scan(&ndata)
	for i := 0; i < ndata; i++{
		fmt.Scan(&info[i].maskapai, &info[i].bandara_asal, &info[i].bandara_tujuan, &info[i].hari, &info[i].waktu_berangkat.jam,&info[i].waktu_berangkat.menit, &info[i].waktu_tiba.jam, &info[i].waktu_tiba.menit)
	}
}
func lamaPerjalanan(info *tabInt, ndata int){
	for i := 0; i < ndata; i++{
		info[i].durasi = (info[i].waktu_tiba.jam - info[i].waktu_berangkat.jam) * 60  + (info[i].waktu_tiba.menit - info[i].waktu_tiba.menit)
	}
}
func cetakArray(info *tabInt, ndata int){
	for i := 0; i < ndata; i++{
		fmt.Println(info[i].maskapai, info[i].bandara_asal, info[i].bandara_tujuan, info[i].hari, info[i].waktu_berangkat.jam, info[i].waktu_berangkat.menit, info[i].waktu_tiba.jam, info[i].waktu_tiba.menit, info[i].durasi)
	}
}