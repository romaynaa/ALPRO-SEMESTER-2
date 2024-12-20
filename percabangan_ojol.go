package main
import "fmt"
func main(){
	var waktu, jam, menit int
	var jarak, tarif float64

	fmt.Scan(&jam, &menit, &jarak)
	waktu = (jam*60) + menit
	if (((waktu >= 300) && waktu <= 540) || ((waktu >= 960 && waktu <= 1140))) && jarak > 0 && jarak <= 10 {
		tarif = 5000 * jarak
		fmt.Println(tarif)
	}else if ((waktu >= 300 && waktu <= 540) || (waktu >= 960 && waktu <= 1140)) && jarak > 10 && jarak <= 20 {
		tarif = 4500 * jarak
		fmt.Println(tarif)
	}else if ((waktu > 1140 && waktu <= 1320) || (waktu > 540 && waktu < 960)) && jarak > 0 && jarak <= 20 {
		tarif = 4000 * jarak
		fmt.Println(tarif)
	}else if waktu > 1320 {
		fmt.Println("maaf, kami tidak bisa melayani pesanan anda")
	}
	
}