package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/now"
)

func main() {

	// var hari = []string{
	// 	"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu",
	// }

	var beginninYear time.Time = now.BeginningOfYear()

	var bulan int16 = 1
	var hari time.Time
	var hariDiSatuBulan = make([]time.Time, 1, 1)
	switch bulan {
	case 1:

		fmt.Println("=======January=======")
		// fmt.Println(hari)
		for i := 0; i <= 31; i++ {
			hari = beginninYear.AddDate(0, 0, i)
			hariDiSatuBulan = append(hariDiSatuBulan, hari)
		}
		fmt.Println(hariDiSatuBulan)
		fmt.Println(hariDiSatuBulan[1].Day)
	case 2:

		fmt.Println("=======Februari=======")
		// fmt.Println(hari)
		for i := 0; i <= 29; i++ {
			hari = beginninYear.AddDate(0, 1, i)
			hariDiSatuBulan = append(hariDiSatuBulan, hari)
		}
		fmt.Println(hariDiSatuBulan)
	case 3:

		fmt.Println("=======Maret=======")
		// fmt.Println(hari)
		for i := 0; i <= 31; i++ {
			hari = beginninYear.AddDate(0, 2, i)
			hariDiSatuBulan = append(hariDiSatuBulan, hari)
		}
		fmt.Println(hariDiSatuBulan)
	case 4:

		fmt.Println("=======April=======")
		// fmt.Println(hari)
		for i := 0; i <= 30; i++ {
			hari = beginninYear.AddDate(0, 3, i)
			hariDiSatuBulan = append(hariDiSatuBulan, hari)
		}
		fmt.Println(hariDiSatuBulan)
	case 5:

		fmt.Println("=======Mei=======")
		// fmt.Println(hari)
		for i := 0; i <= 31; i++ {
			hari = beginninYear.AddDate(0, 4, i)
			hariDiSatuBulan = append(hariDiSatuBulan, hari)
		}
		fmt.Println(hariDiSatuBulan)
	case 6:

		fmt.Println("=======Juni=======")
		// fmt.Println(hari)
		for i := 0; i <= 30; i++ {
			hari = beginninYear.AddDate(0, 5, i)
			hariDiSatuBulan = append(hariDiSatuBulan, hari)
		}
		fmt.Println(hariDiSatuBulan)
	case 7:

		fmt.Println("=======July=======")
		// fmt.Println(hari)
		for i := 0; i <= 31; i++ {
			hari = beginninYear.AddDate(0, 6, i)
			hariDiSatuBulan = append(hariDiSatuBulan, hari)
		}
		fmt.Println(hariDiSatuBulan)
	case 8:

		fmt.Println("=======Agustus=======")
		// fmt.Println(hari)
		for i := 0; i <= 31; i++ {
			hari = beginninYear.AddDate(0, 7, i)
			hariDiSatuBulan = append(hariDiSatuBulan, hari)
		}
		fmt.Println(hariDiSatuBulan)
	case 9:

		fmt.Println("=======September=======")
		// fmt.Println(hari)
		for i := 0; i <= 30; i++ {
			hari = beginninYear.AddDate(0, 8, i)
			hariDiSatuBulan = append(hariDiSatuBulan, hari)
		}
		fmt.Println(hariDiSatuBulan)
	case 10:

		fmt.Println("=======Oktober=======")
		// fmt.Println(hari)
		for i := 0; i <= 31; i++ {
			hari = beginninYear.AddDate(0, 9, i)
			hariDiSatuBulan = append(hariDiSatuBulan, hari)
		}
		fmt.Println(hariDiSatuBulan)
	case 11:

		fmt.Println("=======November=======")
		// fmt.Println(hari)
		for i := 0; i <= 30; i++ {
			hari = beginninYear.AddDate(0, 10, i)
			hariDiSatuBulan = append(hariDiSatuBulan, hari)
		}
		fmt.Println(hariDiSatuBulan)
	case 12:

		fmt.Println("=======Desember=======")
		// fmt.Println(hari)
		for i := 0; i <= 31; i++ {
			hari = beginninYear.AddDate(0, 11, i)
			hariDiSatuBulan = append(hariDiSatuBulan, hari)
		}
		fmt.Println(hariDiSatuBulan)

	default:
		fmt.Println("Tidak ada Bulan")

	}
}
