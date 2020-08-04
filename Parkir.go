package main

// for i := 0; i < 3; i++ {
// 	fmt.Println(rand.Intn(100))

// }
import (
	"fmt"
	"math/rand"
	"time"
)

type PenggunaParkir struct {
	jenisKendaraan string
	platNomor      string
	idParkir       int
	waktuMasuk     int
}

var penggunaParkirArray = make([]PenggunaParkir, 3, 10)

func main() {
	menu()

	var flag bool = false
	var (
		jenisKendaraanParkir string
		platNomorParkir      string
		inputAngka           int
		idParkirKeluar       int
	)

	for !flag {
		fmt.Println("Masukkan Angka pilihan: ")
		fmt.Scan(&inputAngka)

		switch inputAngka {
		case 0:
			menu()
			break

		case 1:
			fmt.Println("Masuk Parkir")

			fmt.Println("Masukkan Jenis Kendaraan")
			fmt.Scan(&jenisKendaraanParkir)

			fmt.Println("Masukkan platNomorParkir")
			fmt.Scan(&platNomorParkir)

			var personParkir = PenggunaParkir{}

			personParkir.jenisKendaraan = jenisKendaraanParkir
			personParkir.platNomor = platNomorParkir
			personParkir.idParkir = idParkirGenerator()

			fmt.Println("Id parkir Anda adalah: ", personParkir.idParkir)
			personParkir.waktuMasuk = time.Now().Second()

			personParkir.parkirMasuk(personParkir.idParkir)
			break
		case 2:
			fmt.Println("Masukkan id Parkir")
			fmt.Scan(&idParkirKeluar)

			parkirKeluar(idParkirKeluar)
			for i := 0; i < len(penggunaParkirArray); i++ {
				if penggunaParkirArray[i].idParkir == idParkirKeluar {
					penggunaParkirArray = append(penggunaParkirArray[:i], penggunaParkirArray[i+1:]...)
				}
			}
			break

		case 3:
			fmt.Println("Keluar dari Menu")
			flag = true
			break
		}

	}

}

func menu() {
	fmt.Println("Selamat Datang di Parkiran G2 Academy")
	fmt.Println("Daftar Menu : ")
	fmt.Println("0. Daftar Menu")
	fmt.Println("1. Masuk Parkir")
	fmt.Println("2. Keluar Parkir")
	fmt.Println("3. Keluar menu")
}

func idParkirGenerator() int {
	var flag bool = false
	var idParkirGenerated int
	for !flag {
		idParkirGenerated = rand.Intn(100)
		for i := 0; i < len(penggunaParkirArray); i++ {
			var person PenggunaParkir = penggunaParkirArray[i]
			if person.idParkir != idParkirGenerated {
				return idParkirGenerated
				flag = true

			}
		}
	}
	return 0

}

func (penggunaParkir PenggunaParkir) parkirMasuk(idParkirMasuk int) {
	var person PenggunaParkir
	for i := 0; i < len(penggunaParkirArray); i++ {
		person = penggunaParkirArray[i]
		if person.idParkir != idParkirMasuk {
			penggunaParkirArray = append(penggunaParkirArray, penggunaParkir)
			fmt.Println("Sudah masuk")
			break
		} else {
			fmt.Println("Parkir udah penuh atau id Parkir sama")
			break
		}
	}
}

func parkirKeluar(idParkirKeluar int) {
	for i := 0; i < len(penggunaParkirArray); i++ {
		var penggunaParkir1 PenggunaParkir = penggunaParkirArray[i]
		if penggunaParkir1.idParkir == idParkirKeluar {
			var selisihWaktu int
			var detikSekarang int = int(time.Now().Second())
			if penggunaParkir1.jenisKendaraan == "mobil" {
				selisihWaktu = detikSekarang - penggunaParkir1.waktuMasuk
				if selisihWaktu == 1 {
					fmt.Println("Biaya Parkir Anda adalah ", " Rp. 5000")
					break
				} else if selisihWaktu > 1 {
					dikurang1Detik := selisihWaktu - 1
					hargaDetikAwal := 5000
					hargaSelanjutnya := 3000 * dikurang1Detik
					totalHarga := hargaDetikAwal + hargaSelanjutnya
					fmt.Println("Biaya Parkir Anda adalah: ", totalHarga)
					break
				} else {
					fmt.Println("Anda belom parkir")
				}
				break
			} else if penggunaParkir1.jenisKendaraan == "motor" {
				selisihWaktu = detikSekarang - penggunaParkir1.waktuMasuk
				if selisihWaktu == 1 {
					fmt.Println("Biaya Parkir Anda adalah ", " Rp. 3000")
					break
				} else if selisihWaktu > 1 {
					dikurang1Detik := selisihWaktu - 1
					hargaDetikAwal := 3000
					hargaSelanjutnya := 2000 * dikurang1Detik
					totalHarga := hargaDetikAwal + hargaSelanjutnya
					fmt.Println("Biaya Parkir Anda adalah: ", totalHarga)
					break
				} else {
					fmt.Println("Anda belom parkir")
				}
				break
			}

		}
	}
}
