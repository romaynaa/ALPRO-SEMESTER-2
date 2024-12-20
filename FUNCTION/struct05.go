package main 

import "fmt"

type waktu struct{
	jam,menit,detik int
}
type kendaraan struct{
	jke, nopol string
	jm, jk waktu
	bea, durasi int
}
func main(){
	var jke, nopol string
	var j1, m1, d1, j2, m2, d2 int
	var k1, k2, k3 kendaraan

	fmt.Scan(&jke,&nopol,&j1,&m1,&d1,&j2,&m2,&d2)
	k1 = kendaraan_baru(jke,nopol,j1,m1,d1,j2,m2,d2)
	fmt.Scan(&jke,&nopol,&j1,&m1,&d1,&j2,&m2,&d2)
	k2 = kendaraan_baru(jke,nopol,j1,m1,d1,j2,m2,d2)
	fmt.Scan(&jke,&nopol,&j1,&m1,&d1,&j2,&m2,&d2)
	k3 = kendaraan_baru(jke,nopol,j1,m1,d1,j2,m2,d2)

	hitungBeaParkir(&k1)
	hitungBeaParkir(&k2)
	hitungBeaParkir(&k3)

	cetak_data(k1)
	cetak_data(k2)
	cetak_data(k3)

}
func kendaraan_baru(j,n string, a,b,c,d,e,f int)kendaraan{
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
func hitungDurasi(k *kendaraan){
	var tot_detik int
	tot_detik = (k.jk.jam - k.jm.jam)*3600 + (k.jk.menit - k.jm.menit)*60 + (k.jk.detik - k.jm.detik)
	k.durasi = tot_detik / 3600
	if tot_detik % 3600 > 0{
		k.durasi++
	}
}
func hitungBeaParkir(k *kendaraan){
	hitungDurasi(k)
	if k.jke == "mobil"{
		k.bea = 6000 + (k.durasi - 1)*3000
	}else if k.jke == "motor"{
		k.bea = 2500 + (k.durasi - 1)*1500
	}
}
func cetak_data(k kendaraan){
	fmt.Printf("Bea parkir mobil %s nopol %s dengan durasi %d jam = Rp %d\n", k.jke,k.nopol,k.durasi,k.bea)
}