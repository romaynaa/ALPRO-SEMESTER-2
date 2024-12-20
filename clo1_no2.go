package main

import "fmt"

func main(){
	var mhs, mhs_br, mhs_lls, mhs_pind, mhs_und, tahun, mhs_tot int
	fmt.Scan(&mhs, &mhs_br, &mhs_lls, &mhs_pind, &mhs_und, &tahun)
	mhs_tot = mhs + (mhs_br - mhs_lls + mhs_pind - mhs_und)* tahun
	fmt.Println(mhs_tot)

}