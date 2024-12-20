package main

import "fmt"

func hitungMenang(g,k int, jm *int){
	if g > k {
		*jm = *jm + 1
	}
}
func hitungDraw(g,k int, jd *int){
	if g == k {
		*jd = *jd + 1
	}
}
func hitungKalah(g,k int, jk *int){
	if g < k {
		*jk = *jk + 1
	}
}
func hitungJumGolKegolanSelisih(g,k int, jg,jkm,jsg *int){
	*jg = *jg + g
	*jkm = *jkm + k
	*jsg = *jg - *jkm
}
func hitungJumPoint(jm,jd,jp *int){
	*jp = *jm * 3 + *jd * 1
}
func main(){
	var n, gol, kegolan, menang, draw, kalah, jumlahg, jumlahk, selisih, total int
	fmt.Scan(&n) 
	for i := 1; i <= n; i++ {
		fmt.Scan(&gol, &kegolan)
		hitungMenang(gol, kegolan, &menang)
		hitungDraw(gol, kegolan, &draw)
		hitungKalah(gol, kegolan, &kalah)
		hitungJumGolKegolanSelisih(gol,kegolan,&jumlahg,&jumlahk,&selisih)
	}
	hitungJumPoint(&menang, &draw, &total)
	fmt.Print(n,menang,draw,kalah,jumlahg,jumlahk,selisih,total)
}
