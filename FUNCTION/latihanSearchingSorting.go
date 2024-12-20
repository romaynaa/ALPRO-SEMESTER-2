//SEARCHING
//SEQUENTIAL SEARCHING
//1. ketemu boolean
var ketemu bool = false
var i = 0
for i < n && !ketemu{
	ketemu = A[i] == x // ketemu = A[i].nama == x
	i++
}
//2. idx integer
var idx, i int 
idx = -1
i = 0
for i < n && idx == -1{
	if A[i].nama = x{
		idx = i
	}
	i++
}
return idx 
// hapus barang
for i := 0; i < n; i++{
	if A[i].nama = nm{
		A[i] = A[n-1]
		n--
	}
}

//BINARY SEARCH
//1. ketemu boolean
var ketemu bool = false
var left, right, mid int
left = 0
right = n - 1
for left <= right && !ketemu{
	mid = (left + right) / 2
	if x > A[mid]{
		right = mid - 1
	}else if x < A[mid]{
		left = mid + 1
	}else{
		ketemu = true
	}
}
return ketemu
//2. idx integer
var left, right, mid int
var idx int = -1
left = 0
right = n - 1
for left <= right && idx == -1{
	mid = (left + right) / 2
	if x > A[mid]{
		right = mid + 1
	}else if x < A[mid]{
		left = mid - 1
	}else{
		idx = mid
	}
}
return idx
//sorting
//1 selection sort
var i, j, idx_min int
var temp nilai
i = 1
for i <= n - 1 {
	idx-min = i - 1
	j = i
	for j < n {
		if A[idx_min].ipk < A[j].ipk{
			idx_min = j
		}
		j++
	}
	temp = A[idx_min]
	A[idx_min] = A[i-1]
	A[i-1] = temp
	i++
}
//2. insertion
// var i,j int
// var temp mahasiswa
// i = 1
// for i <= n - 1 {
// 	j = i
// 	temp = A[j]
// 	for j > 0 && temp.ipk > A[j-1].ipk{
// 		A[j] = A[j-1]
// 		j--
// 	}
// 	A[j] = temp
// 	i++
var i, j int
var temp mahasiswa 
i = 1
for i <= n-1{
	j = i
	temp = A[j]
	for j > 0 && temp.ipk > A[j-1].ipk{
		A[j] = A[j-1]
		j--
	}
	A[j-1] = temp
	i++
}
//hapus barang
for i := 0; i < n; i++{
	if A[i].nama == nama{
		A[i] = A[n-1]
		n--
		return 
	}
}
//SEQUENTIAL SEARCH
var ketemu bool = false
var i int = 0
for i < n && !ketemu {
	ketemu = A[i] = nama
	i++
}
var idx int = -1
var i int = 0
for i < n && idx == -1 {
	if A[i].nama == nama{
		idx = i
	}
	i++
}
return idx
//apabila x tidal ditemukan output -1
var i int = 0
for i < n-1 && A[i] != x{
	i++
}
if A[i] == x{
	return i
}else{
	return -1
}
//BINARY SEARCH
var kr, kn, mid int
var ketemu bool = false
kr = 0
kn = n -1
for kr <= kn && !ketemu{
	mid = (kr+kn)/2
	if A[mid] < x {
		kr = mid + 1
	}else if A[mid] > x {
		kn = mid - 1
	}else{
		ketemu = true
	}
}
return ketemu
//
var kr, kn, mid, idx int
kr = 0
kn = n - 1
idx = -1
for kr <= kn && idx == -1{
	mid = (kr + kn)/2
	if A[mid] > x{
		kn = mid - 1
	}else if A[mid] < x{
		kr = mid + 1
	}else{
		idx = mid
	}
}
return idx 
//hapus data
for i := 0; i < n; i++{
	if A[i].ipk == ipk {
		A[i] = A[n-1]
		n--
		return
	}
}
//SELECTION SORT
var i, j, idx_min int
var temp sekolah
i = 1
for i <= n - 1{
	idx_min = i - 1
	j = i
	for j < n {
		if A[idx_min].ipk < A[j].ipk{
			idx_min = j
		}
		j++
	}
	temp = A[idx_min]
	A[idx_min] = A[i-1]
	A[i-1] = temp
	i++
}
var i, j int
var temp sekolah
i = 1
for i <= n-1 {
	j = i
	temp = A[j]
	for j > 0 && temp.ipk > A[j-1].ipk{
		A[j] = A[j-1]
		j--
	}
	A[j] = temp
	i++
}
//membaca data secara berulang
var A tabInt
var n, x, idx int
fmt.Scan(&x)
n = 0
for x != -1 {
	idx = posisi(A,n,x)
	if idx == -1 {
		A[n].nama = x
		A[n].suara = 1
		n++
	}else{
		A[idx].suara++
	}
	fmt.Scan(&x)
}
//membaca nama dengan spasi
var fname, lname string
for i := 0; i < n && i < NMAX; i++{
	fmt.Scan(&fname, &lname, &A[i].gol; &A[i].assist)
	A[i].name = fname+" "+lname 
}
//hapus data
for i := 0; i < n; i++{
	if A[i].ipk== ipk{
		A[i] = A[n-1]
		n--
		return
	} 
}
//baca data berulang
var A tabGoal 
var n, x, idx int
n = 0
fmt.Scan(&x)
for x != -1 {
	idx = posisi(A,n,x)
	if idx == -1 {
		A[n].nama = x
		A[n].suara = 1
		n++
	}else{
		A[idx].suara++
	}
	fmt.Scan(&x)
}
//sequential searching
var ketemu bool = false
var i int = 0
for i < n && !ketemu {
	ketemu = A[i] == x
	i++
}
return ketemu
//
var idx int = -1
var i int = 0
for i < n && idx == -1 {
	if A[i].ipk == ipk {
	idx = i
	}
	i++
}
return idx
//
var mid int = n / 2
if n % 2 == 0 {
	return float64(A[mid-1]+A[mid])/2
}else{
	return float64(A[mid])
}
func main(){
	fmt.Scan(&x)
	n = 0
	for x != -5313243 && n < NMAX{
		if x == 0 {
			sorting(&A,n)
			fmt.Println(median(A,n))
		}else{
			A[n] = x
			n++
		}
		fmt.Scan(&x)
	}
}
//
var n, x, idx int
fmt.Scan(&x)
n = 0
for x != -1 {
	idx = posisi(A,n,x)
	if idx == -1 {
		A[n].nama = x
		A[n].suara = 1
		n++
	}else{
		A[idx].suara++
	}
	fmt.Scan(&x)
}
//binary search
var ketemu bool = false
var kr, kn, mid int
kr = 0
kn = n - 1
for kr <= kn && !ketemu {
	mid = (kr+kn)/2
	if A[mid] > x{
		kn = mid - 1
	}else if A[mid] < x{
		kr = mid + 1
	}else{
		ketemu = mid
	}
	return ketemu
}
//jika bilangan yang dibaca lebih kecil dari sebelumnya maka pembacaan akan berhenti
var bil int
fmt.Scan(&bil)
for n < NMAX && bil != 0{
	A[n] = bil
	n++
	if bil < A[n-1]{
		bil = 0
	}
}