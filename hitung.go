package main

import (
	"fmt"
	"math"
)

type Angka struct {
	no1 int64
	no2 int64
	no3 int64
}

func main() {

	angka := Angka{}
	angka.no1 = 10
	angka.no2 = 20

	var angkaChange int64 = 10
	change(&angka.no1, &angka.no2, &angka.no3, angkaChange)

	fmt.Println("hasil pangkat ", angka.volumeTabung())
	tambahSepuasnya(1, 2, 3, 4, 5)

}

func tambahSepuasnya(arrayNo ...int) {
	var sum int
	for i := 0; i < len(arrayNo); i++ {
		sum += arrayNo[i]
	}
	fmt.Println("total:", sum)
}

func change(original1 *int64, original2 *int64, original3 *int64, value int64) {
	*original1 = *original1
	*original2 = *original2
	*original3 = *original3
}

func (angka Angka) perkalian() int64 {
	return angka.no1 * angka.no2 * angka.no3
}

func (angka Angka) bagi() int64 {
	return angka.no1 / angka.no2 / angka.no3
}

func (angka Angka) tambah() int64 {
	return angka.no1 + angka.no2 + angka.no3
}

func (angka Angka) kurang() int64 {
	return angka.no1 - angka.no2 - angka.no3
}

func (angka Angka) akar() float64 {
	akar1 := math.Sqrt(float64(angka.no1))

	return akar1

}

func (angka Angka) pangkat() float64 {
	pangkat1 := math.Pow(float64(angka.no1), 2)

	return pangkat1
}

func (angka Angka) luasPersegi() int64 {
	return angka.no1 * angka.no2
}

func (angka Angka) luasLingkaran() float64 {
	var radius float64 = float64(angka.no1)

	return math.Phi * radius * radius
}

func (angka Angka) volumeTabung() float64 {
	var radius float64 = float64(angka.no1)
	var panjang float64 = float64(angka.no2)

	return math.Phi * radius * radius * panjang
}

func (angka Angka) volumeBalok() int64 {
	return angka.no1 * angka.no2 * angka.no3
}

func (angka Angka) volumePrisma() int64 {
	return angka.no1 * angka.no2 * angka.no3 / 2
}
