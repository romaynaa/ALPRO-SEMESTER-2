package main 

import "fmt" 

const NMAX int = 10

type books struct{
	boookID, title, genre, name string
	terbit int
}
type tabBooks [NMAX]books

func main(){
	var data tabBooks
	var nData int
	var boookID stirng

	fmt.Scan(&data,&nData)

}
func baca_data(A *tabBooks, n *int){
	if *n > NMAX{
		*n = NMAX
	}
	for i := 0; i < *n; i++{
		fmt.Scan(&A[i].boookID,&A[i].title,&A[i].genre,&A[i].name,&A[i].terbit)
	}
}
func cetak_Data(A tabBooks, n int){
	for i := 0; i < n; i++{
		fmt.Println(A[i].boookID,A[i].title,A[i].genre,A[i].name,A[i].terbit)
	}
}
func edit_data(A tabBooks, n int, bID string){

}
func find_data(A *tabBooks, n int, bID string)int{
	var idx int = -1
	var i int = 0
	for i < n && idx == -1{
		A[i].boookID == bID
		idx = i
	}
}