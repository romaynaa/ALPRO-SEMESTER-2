package main 

import "fmt"

type mobil struct{
	nama string
	tahun, v int
}
func main(){
	var m1,m2,m3 mobil
	fmt.Scan(&m1.nama, &m1.tahun, &m1.v)
	fmt.Scan(&m2.nama, &m2.tahun, &m2.v)
	fmt.Scan(&m3.nama, &m3.tahun, &m3.v)
	hitungRerata(m1,m2,m3)
	cetakData(m1,m2,m3)
}
func hitungRerata(m1, m2, m3 mobil)float64{
	var rata float64
	rata = float64(m1.v + m2.v + m3.v)/3
	return rata
}
func cetakData(m1,m2,m3 mobil){
	fmt.Printf("Rata-rata kecepatan mobil %s %d", m1.nama, m1.tahun)
	fmt.Printf(", mobil %s %d", m2.nama, m2.tahun)
	fmt.Printf(", dan mobil %s %d", m3.nama, m3.tahun)
	fmt.Printf(",adalah %.2f", hitungRerata(m1,m2,m3))
}