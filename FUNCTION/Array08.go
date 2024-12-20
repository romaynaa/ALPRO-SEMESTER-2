package main 

import "fmt"

const NMAX int = 1000

type wisuda struct{
	nama, nim string 
	eprt, semester int
	ipk float64
}
type tabInt[NMAX] wisuda 

func main(){
	var data tabInt
	var ndata int
	bacaData(&data, &ndata)
	cetakData(data, ndata)
	


}
func bacaData(d *tabInt, n* int){
	var nma, nm string
	var ep, sem int
	var ip float64
	var i int = 0
	fmt.Scanln(&nma,&nm,&ep,&sem,&ip)
	for nma != "none" && i < NMAX{
		d[i].nama = nma
		d[i].nim = nm
		d[i].eprt = ep
		d[i].semester = sem
		d[i].ipk = ip
		i++
		fmt.Scanln(&nma,&nm,&ep,&sem,&ip)
	}
	*n = i
}
func catakData(d tabInt, n int){
	fmt.Println("Data Wisudawan:")
	for i := 0; i < n; i++{
		fmt.Printf("%s %s %d %d %f\n", d[i].nama, d[i].nim, d[i].eprt, d[i].semester, d[i].ipk)
	}
}
func maxEprt(d tabInt, n int)int{
	var k int = 1
	var idxMax int = 0
	for k < n {
		if d[idxMax] < d[k].eprt{
			idxMax = k
		}
		k += 1
	}
	return idxMax
}
func minIpk(d tabInt, n int)int{
	var k int  = 1
	var idxMin int = 0
	for k < n {
		if d[idxMin] > d[k].ipk{
			idxMin = k
		}
		k += 1
	}
	return idxMin
}
func rataArray(d tabInt, n int)float64{
	var jumlah int
	var k int = 1
	for k = 0; k < n; k++{
		jumlah = jumlah + d[k].semester
	}
	return jumlah/n
}