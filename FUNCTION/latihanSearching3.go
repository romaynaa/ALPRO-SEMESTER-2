package main 

import "fmt"

type partai struct{
	nama, suara int
}
const NMAX int = 1000000
type tabPartai [NMAX]partai

func main(){
	var data tabPartai
	var nData, x, idx int

	// nData = 0
	// fmt.Scan(&x)
	// for x != -1 {
	// 	idx = posisi(data, nData, x)
	// 	if idx == -1 {
	// 		data[nData].nama = x
	// 		data[nData].suara = 1
	// 		nData++
	// 	}else{
	// 		data[idx].suara++
	// 	}
	// 	fmt.Scan(&x)
	// }
	nData = 0
	fmt.Scan(&x)
	for x != -1{
		idx = posisi(data, nData, x)
		if idx == -1{
			data[idx].nama = x
			data[idx].suara = 1
			nData++
		}else{
			data[idx].suara++
		}
		fmt.Scan(&x)
	} 

	// var i, j int
	// var temp partai
	// i = 1
	// for i <= nData - 1{
	// 	j = i
	// 	temp = data[j]
	// 	for j > 0 && temp.suara > data[j-1].suara{
	// 		data[j] = data[j-1]
	// 		j--
	// 	}
	// 	data[j] = temp
	// 	i++
	// }
	//insertion sort
	var i, j int
	var temp partai
	i = 1
	for i <= nData-1{
		j = i
		temp = data[j] 
		for j > 0 && temp.suara > data[j-1].suara{
			data[j] = data[j-1]
			j--
		}
		data[j] = temp
		i++
	}
	//menampilkan
	for i := 0; i < nData; i++{
		fmt.Printf("%v(%v) ", data[i].nama, data[i].suara)
	}
}
func posisi(A tabPartai, n int, nama int)int{
	// var idx, i int
	// idx = -1
	// i = 0
	// for i < n && idx == -1{
	// 	if A[i].nama == nama{
	// 		idx = i
	// 	}
	// 	i++
	// }
	// return idx
	var idx, i int
	idx = -1
	i = 0
	for i < n && idx == -1{
		if A[i].nama == nama{
			idx = i
		}
		i++
	}
	return idx
}