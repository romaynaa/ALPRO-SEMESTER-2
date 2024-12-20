package main 

import (
	"fmt"
	"os"
	"os/exec"
)
func intro(){
	clear_screen()
	fmt.Println("Selamat Datang")
}
func bye(){
	clear_screen()
	fmt.Println("Sampai Jumpa")
}
func clear_screen(){
	c:= exec.Command("Clear")
	c.Stdout = os.Stdout
	c.Run()
}
func menu_utama(pilih int){
	fmt.Println("-----------------------")
	fmt.Println("        M E N U        ")
	fmt.Println("-----------------------")
	fmt.Println("1. Jumlah bilangan asli")
	fmt.Println("2. Faktorial")
	fmt.Println("3. Fibonacci")
	fmt.Println("4. Exit")
	fmt.Println("-----------------------")
	fmt.Print("Pilih (1/2/3/4) : ")
	fmt.Scan(&pilih)
}
func menu_jumlah_bilangan_asli(){
	clear_screen()
	fmt.Println("--------------------------")
	fmt.Println("          M E N U         ")
	fmt.Println("--------------------------")
	fmt.Println("1. Jumlah Bilangan Asli   ")
	fmt.Println("2. Exit                   ")
	fmt.Println("--------------------------")
}
func jumlah_bilangan_asli(){
	var pilih, N, jumlah int
	var lagi string
	for {
		menu_jumlah_bilangan_asli()
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			for {
				fmt.Print("Masukkan bilangan asli N: ")
				fmt.Scan(&N)
				jumlah = jumlah_bilangan_asli_rekursif(N)
				fmt.Printf("jumlah %d bilangan asli pertama %d\n", N, jumlah)
				fmt.Print("Lagi (Y/T): ")
				fmt.Scan(&lagi)
				if lagi != "Y"{
					break
				}
			}
			clear_screen()
		}else if pilih == 2 {
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
	clear_screen()
	fmt.Println("--------------------------")
	fmt.Println("          M E N U         ")
	fmt.Println("--------------------------")
	fmt.Println("1. Faktorial  ")
	fmt.Println("2. Exit                   ")
	fmt.Println("--------------------------")
}
func faktorial(){
	var N,pilih,suku int
	var lagi string
	for {
		menu_faktorial()
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&pilih)
		if pilih == 1{
			for {
				fmt.Print("Masukkan bilangan asli N: ")
				fmt.Scan(&N)
				fmt.Printf("Nilai faktorial %d: %d\n", N, suku)
				fmt.Print("Lagi (Y/T: ")
				fmt.Scan(&lagi)
				if lagi != "Y" {
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
func faktorial_rekursif(n int)int{
	if n == 1{
		return 1
	}
	return n * faktorial_rekursif(n-1)
}

func menu_fibonacci(){
	clear_screen()
	fmt.Println("--------------------------")
	fmt.Println("          M E N U         ")
	fmt.Println("--------------------------")
	fmt.Println("1. Fibonacci  ")
	fmt.Println("2. Exit                   ")
	fmt.Println("--------------------------")
}
func fibonacci(){
	var N, pilih, suku int
	var lagi int
	for{
		menu_fibonacci()
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&pilih)
		if pilih == 1{
			for {
				fmt.Print("Masukan bilangan asli N: ")
				fmt.Scan(&N)
				suku == fibonacci_rekursif(N)
				fmt.Printf("Suku fibonacci ke-%d: %d\n: ")
				fmt.Print("lagi (Y/T): ")
				fmt.Scan(&lagi)
				if lagi != "Y"{
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
func fibonacci_rekursif(n int)int{
	if n == 1{
		return 0
	}else if n == 2{
		return 1
	}
	return fibonacci_rekursif(n-1)+fibonacci_rekursif(n-2)
}

func main(){
	var pilih int
	intro()
	for {
		menu_utama(&pilih)
		switch pilih{
		case 1:
			jumlah_bilangan_asli()
		case 2:
			faktorial()
		case 3:
			fibonacci()
		default:
			clear_screen()
		}
		if pilih == 4{
			break
		}
	}
	bye()
}