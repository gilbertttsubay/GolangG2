package main

// for i := 0; i < 3; i++ {
// 	fmt.Println(rand.Intn(100))

// }
import (
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
)

type PenggunaParkir struct {
	jenisKendaraan string
	platNomor      string
	idParkir       int
	waktuMasuk     int
}

type idParkirKeClientStruct struct{
	idParkirKeClient int
}

var penggunaParkirArray = make([]PenggunaParkir, 3, 10)


//struct untuk mengambil data inputan jenis kendaraan dan plat nomor dari client
type ParkerDariClient struct {
	Jenis_kendaraan string `json: "jenis_kendaraan"`
	Plat_nomor string `json: "plat_nomor"`
}

func main() {
	http.HandleFunc("/id_parkir", parkirMasukDariClient)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
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

func (penggunaParkir PenggunaParkir) parkirMasuk(idParkirMasuk int) string {
	var person PenggunaParkir
	for i := 0; i < len(penggunaParkirArray); i++ {
		person = penggunaParkirArray[i]
		if person.idParkir != idParkirMasuk {
			penggunaParkirArray = append(penggunaParkirArray, penggunaParkir)
			return "Sudah masuk"
			break
		} else {
			fmt.Println("Parkir udah penuh atau id Parkir sama")
			break
		}
	}
	return "Parkir sudah penuh"
}

func parkirKeluar(idParkirKeluar int) int64 {
	for i := 0; i < len(penggunaParkirArray); i++ {
		var penggunaParkir1 PenggunaParkir = penggunaParkirArray[i]
		if penggunaParkir1.idParkir == idParkirKeluar {
			var selisihWaktu int
			var detikSekarang int = int(time.Now().Second())
			if penggunaParkir1.jenisKendaraan == "mobil" {
				selisihWaktu = detikSekarang - penggunaParkir1.waktuMasuk
				if selisihWaktu == 1 {
					return 5000
					fmt.Println("Biaya Parkir Anda adalah ", " Rp. 5000")
					break
				} else if selisihWaktu > 1 {
					dikurang1Detik := selisihWaktu - 1
					hargaDetikAwal := 5000
					hargaSelanjutnya := 3000 * dikurang1Detik
					totalHarga := hargaDetikAwal + hargaSelanjutnya
					return int64(totalHarga)
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
					return int64(totalHarga)
					break
				} else {
					fmt.Println("Anda belom parkir")
				}
				break
			}

		}
	
	}
	return 0
}

func parkirMasukDariClient(res http.ResponseWriter, req *http.Request) {
	//response dari client
	responseDariClient, error := ioutil.ReadAll(req.Body)

	//penampung json
	var parkirDariClient ParkerDariClient

	//penampung dari penampung json
	var penggunaParkir PenggunaParkir

	//"Membuka format json ke bentukannnya ParkerDariClient"
	error = json.Unmarshal(responseDariClient, &parkirDariClient)
	if error != nil {
		log.Fatalf("Tidak bisa meng-unmarshal response dari client dikarenakan ada error: ", error)
	}

	//meng-generate id parkir
	idGenerated := idParkirGenerator()

	//mendefinisikan parkir dari client ke pengguna parkir
	if idGenerated != 0 {
		penggunaParkir.idParkir = idGenerated
		penggunaParkir.jenisKendaraan = parkirDariClient.Jenis_kendaraan
		penggunaParkir.platNomor = parkirDariClient.Plat_nomor
		penggunaParkir.waktuMasuk = time.Now().Second()

		penggunaParkirArray = append(penggunaParkirArray, penggunaParkir)

		//mengirim response id ke client
		responseKeClient := idGenerated

		var idYangAkanDikirim = idParkirKeClientStruct{responseKeClient}

		jsonIdYangDikirim, error :=json.Marshal(idYangAkanDikirim)
		if error != nil {
			log.Fatalf("Tidak bisa mengubah id yang dikirim ke dalam bentuk Json dikarenakan: ", error)
		}

		res.Write(jsonIdYangDikirim)
	}




}
