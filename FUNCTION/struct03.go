package main

import "fmt"

type pegawai struct{
	nama string
	gaji, masa int
	bonus float64
}
func main(){
	var d pegawai
	fmt.Scan(&d.nama,&d.gaji,&d.masa)
	hitung_bonus(&d)
	fmt.Println("Pegawai dengan nama",d.nama,"mendapatkan bonus sebesar Rp ",d.gaji)

}
func hitung_bonus(d *pegawai){
	if d.masa >= 10{
		d.gaji = d.gaji * 3/2
	}else if d.masa <= 10 && d.masa >= 5{
		d.gaji = d.gaji * 3/4
	}else{
		d.gaji = d.gaji * 1/2
	}
}