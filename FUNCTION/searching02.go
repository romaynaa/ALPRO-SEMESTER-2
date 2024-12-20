//SEQUENTIAL SEARCH

/*
kamus
	ketemu : boolean
	i : integer
algoritma 
	ketemu <- false
	i = 0
	while !ketemu and i < n do
		ketemu <- T[i] == x
		i++
	endwhile
	return ketemu
endfunction
*/
/*
kamus
	idx : integer
	i : integer
algoritma
	idx <- -1
	i <- 0
	for idx == -1 and i < n do
		if T[i] == x{
			idx <- i
		}
	endwhile
	return idx
endfunction
*/
// var idx_min, temp int
// var i int = 1
// var j int
// for i <= n - 1 {
// 	idx_min = i - 1
// 	j = i
// 	for j < n{
// 		if A[idx_min] > A[j] {
// 			idx_min = j
// 		}
// 		j++
// 	}
// 	temp = A[idx_min]
// 	A[idx_min] = A[i-1]
// 	A[i-1] = temp
// 	i++
// }
insertion sort
for i := 0; i < n; i++{
	temp := A[i]
	j = i - 1
	for j >= 0 && A[j].durasi > temp.durasi{
		A[j+1] = A[j]
		j--
	}
	A[j+1] = temp
}
selection sort
for i := 0; i < n; i++{
	idx_min := i
	for j := i + 1; j < n; j++{
		if A[j].rating > A[idx_min].rating{
			idx_min = j
		}
	}
	A[i],A[idx_min] = A[idx_min],A[i]
}
sequencial search
var xnilai int
var found bool = false
var i int = 0
for i < n && !found{
	found = A[i] == xnilai
	i++
}
sequencial search apabila x tidak ditemukan x = -1{
var found int = -1
var i = 0
var xnilai int
for i < n; && found == -1{
		if A[i] == xnilai{
			found = i
			}
	i++
	}
	return found
}

binary search
var x int
var found bool = false
var kr, kn, med int
kr = 0
kn = n-1
for kr <= kn && !found {
	med = (kr+kn)/2
	if A[med] < x{
		kr = med+1
	}else if A[med] > x{
		kn = med
	}else{
		found = true	
	}
}
binary search apabila x tidak ditemukan x = -1
var found int = -1
var kr int = 0
var kn int = n-1
var x int
for kr <= kn && found == -1{
	med = (kr+kn)/2
	if x > A[med]{
		kn = med -1
	}else if x < A[med]{
		kr = med + 1
	}else{
		found = med
	}
	return found
}
