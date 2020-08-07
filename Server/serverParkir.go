package main

import (
	"fmt"
	"TugasParkir/gRPC"
	"log"
	"net"
	"context"
	"time"
	"math/rand"

	"google.golang.org/grpc"
)


var penggunaParkirArray = make([]PenggunaParkir, 3, 10)

type PenggunaParkir struct {
	jenisKendaraan string
	platNomor      string
	idParkir       int64
	waktuMasuk     int
}

func(*server) MasukParkir(ctx context.Context, req *tugasparkirpb.MasukParkirRequest) (*tugasparkirpb.MasukParkirResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	idGenerated:= idParkirGenerator()
	if idGenerated != 0 {
		var personParker = PenggunaParkir{}
		personParker.jenisKendaraan = req.GetMasukParkir().GetJenisKendaraan()
		personParker.platNomor= req.GetMasukParkir().GetPelatNomor()
		personParker.idParkir = idGenerated
		personParker.waktuMasuk = time.Now().Second()
		var statusParkir = personParker.parkirMasuk(idGenerated)
		if statusParkir == "Bisa masuk"{
			res := &tugasparkirpb.MasukParkirResponse{
				GeneratedId: idGenerated,
			}
			return res, nil
		}
	
	}
	res := &tugasparkirpb.MasukParkirResponse{
		GeneratedId: 0,
	}

	return res, nil

}

func(*server) KeluarParkir(ctx context.Context, req *tugasparkirpb.KeluarParkirRequest) (*tugasparkirpb.KeluarParkirResponse, error){
	fmt.Printf("Greet function was invoked with %v\n", req)
	
	idparkir:= req.GetKeluarParkir().GetIdKeluar()
	totalHarga := parkirKeluar(idparkir)


	res:= &tugasparkirpb.KeluarParkirResponse{
		TotalBiaya: totalHarga,
	}
	return res, nil
}
func (penggunaParkir PenggunaParkir) parkirMasuk(idParkirMasuk int64) string {
	var person PenggunaParkir
	for i := 0; i < len(penggunaParkirArray); i++ {
		person = penggunaParkirArray[i]
		if person.idParkir != idParkirMasuk {
			penggunaParkirArray = append(penggunaParkirArray, penggunaParkir)
			return "Bisa masuk"
			break
		} else {
			return "Parkir udah penuh atau id Parkir sama"
			break
		}
	}
	return "Parkir udah penuh"
}

func idParkirGenerator() int64 {
	var flag bool = false
	var idParkirGenerated int64
	for !flag {
		idParkirGenerated = int64(rand.Intn(100))
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



func parkirKeluar(idParkirKeluar int64) int64  {
	for i := 0; i < len(penggunaParkirArray); i++ {
		var penggunaParkir1 PenggunaParkir = penggunaParkirArray[i]
		if penggunaParkir1.idParkir == idParkirKeluar {
			var selisihWaktu int
			var detikSekarang int = int(time.Now().Second())
			if penggunaParkir1.jenisKendaraan == "mobil" {
				selisihWaktu = detikSekarang - penggunaParkir1.waktuMasuk
				if selisihWaktu == 1 {
					fmt.Println("Biaya Parkir Anda adalah ", " Rp. 5000")
					return 5000
					break
				} else if selisihWaktu > 1 {
					dikurang1Detik := selisihWaktu - 1
					hargaDetikAwal := 5000
					hargaSelanjutnya := 3000 * dikurang1Detik
					totalHarga := int64(hargaDetikAwal + hargaSelanjutnya)
					fmt.Println("Biaya Parkir Anda adalah: ", totalHarga)
					return totalHarga
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
					totalHarga := int64(hargaDetikAwal + hargaSelanjutnya)
					fmt.Println("Biaya Parkir Anda adalah: ", totalHarga)
					return totalHarga
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

type server struct{}
type ServerKeluar struct{}

func main() {
	fmt.Println("Hello")

	lis, error := net.Listen("tcp", "0.0.0.0:50051")

	if error != nil {
		log.Fatal("Failed to listen: %v", error)
	}

	s := grpc.NewServer()

	tugasparkirpb.RegisterParkirServiceServer(s, &server{})

	if error := s.Serve(lis); error != nil {
		log.Fatalf("failed to serve: %v", error)
	}

}
