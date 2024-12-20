// // // // // // package main

// // // // // // import "fmt"

// // // // // // func barisan(n int)int{
// // // // // // 	if n <= 3{
// // // // // // 		return 1
// // // // // // 	}
// // // // // // 	return barisan(n-3)+barisan(n-2)+barisan(n-1)
// // // // // // }
// // // // // // func main(){
// // // // // // 	var n int
// // // // // // 	fmt.Scan(&n)
// // // // // // 	fmt.Println(barisan(n))
// // // // // // }
// // // // // package main

// // // // // import "fmt"

// // // // // func rerata(n, jumlah, i int, rata float64){
// // // // // 	if i > n{
// // // // // 	return
// // // // // 	}
// // // // // 	jumlah += i
// // // // // 	rata = float64(jumlah) / float64(i)
// // // // // 	fmt.Println(jumlah, rata)
// // // // // 	rerata(n, jumlah, i+1, rata)
// // // // // }
// // // // // func main(){
// // // // // 	var n int
// // // // // 	fmt.Scan(&n)
// // // // // 	rerata(n,0,1,0)
// // // // // }
// // // // package main

// // // // import "fmt"

// // // // func gcd(a, b int)int{
// // // // 	if b == 0{
// // // // 		return a
// // // // 	}
// // // // 	return gcd(b, a%b)
// // // // }
// // // // func main(){
// // // // 	var a,b int
// // // // 	fmt.Scan(&a,&b)
// // // // 	fmt.Println(gcd(a,b))
// // // // }
// // // package main

// // // import "fmt"

// // // func gcd(a, b int)int{
// // // 	if b == 0{
// // // 		return a
// // // 	}
// // // 	return gcd(b, a%b)
// // // }
// // // func lcm(a,b int)int{
// // // 	return a * b / gcd(a,b)
// // // }
// // // func main(){
// // // 	var a,b int
// // // 	fmt.Scan(&a,&b)
// // // 	fmt.Println(lcm(a,b))
// // // }
// // package main

// // import ("fmt"
// // 		"math")

// // func panjangX(R,T float64)float64{
// // 	X := R * math.Cos(T)
// // 	return X
// // }
// // func panjangY(R,T float64)float64{
// // 	Y := R * math.Sin(T)
// // 	return Y
// // }
// // func main(){
// // 	var R, T, RX, RY float64
// // 	fmt.Scan(&R, &T)
// // 	T = T * math.Pi/180.0
// // 	RX = panjangX(R,T)
// // 	RY = panjangY(R,T)
// // 	fmt.Printf("%10.2f %10.2f",RX,RY)

// // }\
// package main

// import "fmt"

//	func faktorial(n int)int{
//		if n == 1 || n == 0{
//			return 1
//		}else{
//		return n * faktorial(n-1)
//		}
//	}
//
//	func main(){
//		var a,b,c,d int
//		var p1,p2,c1,c2 int
//		fmt.Scan(&a,&b,&c,&d)
//		if a > c{
//			p1 = faktorial(a)/faktorial(a-c)
//			c1 = faktorial(a)/(faktorial(c)*faktorial(a-c))
//		}else if b > d{
//			p1 = faktorial(b)/faktorial(b-d)
//			c1 = faktorial(b)/(faktorial(d)*faktorial(b-d))
//		}
//		fmt.Println(p1,c1,p2,c2)
//	}
// package main

// import "fmt"

// func faktorBilangan(n, i int)int{
// 	if i > n{
// 		return 1
// 	}else{
// 	if n % i == 0{
// 	fmt.Print(i)
// 	faktorBilangan(n,i+1)
// 		}
// 	}
// 	return i
// }
// func main(){
// 	var n int
// 	fmt.Scan(&n)
// 	faktorBilangan(n,1)
// }
package main

import "fmt"

func pangkat(x,y int)int{
	if y == 0{
		return 1
	}
	return x * pangkat(x, y-1)
}
func main(){
	var x, y int
	fmt.Scan(&x,&y)
	fmt.Println(pangkat(x,y))
}