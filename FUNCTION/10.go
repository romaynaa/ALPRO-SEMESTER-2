package main

import (
	"fmt"
	"os"
	"os/exec"
)
func intro(){
	clear_screen()
	fmt.Println("Selamat datang")
}
func bye(){
	clear_screen()
	fmt.Println("Sampai Jumpa")
}
func clear_screen(){
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
func menu_utama(p *int){
	fmt.Println("-----------------------")
	fmt.Println("        M E N U        ")
	fmt.Println("-----------------------")
	fmt.Println("1. Jumlah bilangan asli")
	fmt.Println("2. Faktorial")
	fmt.Println("3. Fibonacci")
	fmt.Println("4. Exit")
	fmt.Println("-----------------------")
	fmt.Print("Pilih (1/2/3/4)? ")
	fmt.Scan(p)
}
func menu_jumlah_bilangan_asli(){
	fmt.Println("-----------------------")
	fmt.Println("        M E N U        ")
	fmt.Println("-----------------------")
	fmt.Println("1. Jumlah bilangan asli")
	fmt.Println("2. Exit")
	fmt.Println("-----------------------")
}
func jumlah_bilangan_asli(){
	var lagi string
	var n, jumlah, pilih int
	for {
		menu_jumlah_bilangan_asli()
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&pilih)
		if pilih == 1{
			for {
				fmt.Print("Masukkan bilangan asli n: ")
				fmt.Scan(&n)
				jumlah = jumlah_bilangan_asli_rekursif(n)
				fmt.Printf("Jumlah bilangan asli %d:%d\n ",n, jumlah)
				fmt.Print("Lagi (Y/T)? ")
				fmt.Scan(&lagi)
				if lagi == "T"{
					break
				}
			}
			clear_screen()
		}else if pilih == 2{
			clear_screen()
			break
		}
	}
}
func jumlah_bilangan_asli_rekursif(n int)int{
	if n == 1{
		return 1
	}
	return n + jumlah_bilangan_asli_rekursif(n-1)
}

func menu_faktorial(){
	fmt.Println("-----------------------")
	fmt.Println("        M E N U        ")
	fmt.Println("-----------------------")
	fmt.Println("1. Faktorial")
	fmt.Println("2. Exit")
	fmt.Println("-----------------------")
}
func faktorial(){
	var n, jumlah, p int
	var lagi string
	
	for {
		menu_faktorial()
		fmt.Scan(&p)
		fmt.Print("Pilih (1/2): ")
		if p == 1{
			for {
				fmt.Print("Masukkan bilangan asli n: ")
				fmt.Scan(&n)
				jumlah = faktorial_rekursif(n)
				fmt.Printf("Nilai faktorial %d:%d\n",n, jumlah)
				fmt.Print("Lagi (Y/T)? ")
				fmt.Scan(&lagi)
				if lagi == "T" {
					break
				}
			}
		clear_screen()
		}else if p == 2{
			clear_screen()
			break
		}
	}
}
func faktorial_rekursif(n int)int{
	if n == 1{
		return 1
	}
	return n * faktorial_rekursif(n-1)
}

func menu_fibonacci(){
	fmt.Println("-----------------------")
	fmt.Println("        M E N U        ")
	fmt.Println("-----------------------")
	fmt.Println("1. Fibonacci")
	fmt.Println("2. Exit")
	fmt.Println("-----------------------")
}
func fibonacci(){
	var n, jumlah, p int
	var lagi string
	for {
		menu_fibonacci()
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&p)
		if p == 1{
			for {
			fmt.Print("Masukkan bilangan asli n: ")
			fmt.Scan(&n)
			jumlah = fibonacci_rekursif(n)
			fmt.Printf("Suku fibonacci ke-%d:%d\n",n, jumlah)
			fmt.Print("Lagi (Y/T)? ")
			fmt.Scan(&lagi)
			if lagi == "T"{
				break
			}
		}
		clear_screen()
		}else if p == 2{
			clear_screen()
			break
		}
	}
}
func fibonacci_rekursif(n int)int{
	if n == 1{
		return 0
	}else if n == 2{
		return 1
	}
	return n + fibonacci_rekursif(n-1) + fibonacci_rekursif(n-2)
}

func main(){
	var pilih int
	intro()
	for {
		menu_utama(&pilih)
		switch pilih{
		case 1: jumlah_bilangan_asli()
		case 2: faktorial()
		case 3: fibonacci()
		default :
			clear_screen()
		}
		if pilih == 4{
			break
		}
	}
	bye()
}