package main

import "fmt"

type waktu struct{
	jam, menit, detik int
}
type kendaraan struct{
	jke, nopol string
	jm, jk waktu
	bea, durasi int
}
func main(){
	var k01,k02,k03 kendaraan
	var jenis, nopol string
	var j1,m1,d1,j2,m2,d2 int

	fmt.Scan(&jenis,&nopol,&j1,&m1,&d1,&j2,&m2,&d2)
	//membuat kendaraan k01
	k01 = kendaraan_baru(jenis,nopol,j1,m1,d1,j2,m2,d2)

	fmt.Scan(&jenis,&nopol,&j1,&m1,&d1,&j2,&m2,&d2)
	//membuat kendaraan k02
	k02 =  kendaraan_baru(jenis,nopol,j1,m1,d1,j2,m2,d2)

	fmt.Scan(&jenis,&nopol,&j1,&m1,&d1,&j2,&m2,&d2)
	//membuat kendaraan k03
	k03 =  kendaraan_baru(jenis,nopol,j1,m1,d1,j2,m2,d2)

	//hitung bea parkir 3 kendaraan
	bea_parkir(&k01)
	bea_parkir(&k02)
	bea_parkir(&k03)
	//cetak data parkir per kendaraan
	cetak_data(k01)
	cetak_data(k02)
	cetak_data(k03)
}
func kendaraan_baru(j, n string, a,b,c,d,e,f int)kendaraan{
	//mengembalikan nilai kendaraan
	var k kendaraan 
	k.jke = j
	k.nopol = n
	k.jm.jam = a
	k.jm.menit = b
	k.jm.detik = c
	k.jk.jam = d
	k.jk.menit = e
	k.jk.detik = f
	return k
}
func durasi_parkir(k *kendaraan){
	var tot_detik int
	tot_detik = (k.jk.jam-k.jm.jam)*3600 + (k.jk.menit-k.jm.menit)*60 + (k.jk.detik-k.jm.detik)
	k.durasi = tot_detik/3600
	if tot_detik % 3600 > 0{
		k.durasi++
	}
}
func bea_parkir(k *kendaraan){
	durasi_parkir(k)
	if k.jke == "mobil"{
		k.bea = 6000 + (k.durasi - 1)*3000
	}else if k.jke == "motor"{
		k.bea = 2500 + (k.durasi - 1)*1500
	}
}
func cetak_data(k kendaraan){
	fmt.Printf("bea parkir %s nopol %s dengan durasi %d jam: Rp %d\n", k.jke, k.nopol, k.durasi, k.bea)

}