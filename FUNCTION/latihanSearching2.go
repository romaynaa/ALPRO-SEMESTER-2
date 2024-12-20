package main 

import "fmt"

type bola struct{
	nama string
	goal, assist int
}
const NMAX int = 10
type tabGoal [NMAX]bola

func main(){
	var data tabGoal
	var ndata int

	bacaData(&data, &ndata)
	
	urutMenurun(&data, ndata)
	// fmt.Println("Data setelah diurutkan: ")
	// cetakData(data, ndata)
}
func bacaData(A *tabGoal, n *int){
	fmt.Scan(n)
	if *n > NMAX{
		*n = NMAX
	}
	var fname, lname string
	for i := 0; i < *n; i++{
		fmt.Scan(&fname, &lname,&A[i].goal, &A[i].assist)
		A[i].nama = fname+" "+lname
	}
}
// func cetakData(A tabGoal, n int){
	
// 	for i := 0; i < n; i++{
// 		fmt.Println(A[i].nama, A[i].goal, A[i].assist)
// 	}
// }
func urutMenurun(A *tabGoal, n int){
	/*var i, j int
	var temp bola
	i = 1
	for i <= n-1 {
		j = i
		temp = A[j]
		for (j > 0 && temp.goal > A[j].goal) || (temp.goal == A[j].goal && temp.assist > A[j].assist){
			A[j] = A[j-1]
			j--
		}
		A[j] = temp
		i++
	}
	*/
	var i, j, idx_min int
	var temp bola
	i = 1
	for i <= n - 1{
		idx_min = i - 1
		j = i
		for j < n {
			if A[idx_min].goal < A[j].goal || (A[idx_min].goal == A[j].goal && A[idx_min].assist <= A[j].assist){
				idx_min = j
			}
			j++
		}
		temp = A[idx_min]
		A[idx_min] = A[i-1]
		A[i-1] = temp
		i++
	}
	//menampilkan
	for i := 0; i < n; i++{
		fmt.Println(A[i].nama, A[i].goal, A[i].assist)
	}
}