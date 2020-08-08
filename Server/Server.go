package main
import(
	"Parkir_gRPC_Server_Streaming_rpc/Proto"
	"math/rand"
	"fmt"
	"time"
	"log"
	"google.golang.org/grpc"
	"net"
)

type PenggunaParkir struct {
	jenisKendaraan string
	platNomor      string
	idParkir       int64
	waktuMasuk     int
}

type PenggunaParkirArray struct {
	penggunaParkirArray []PenggunaParkir
}

var penggunaParkirArray = make([]PenggunaParkir, 3, 10)


type server struct {}





func main(){
	fmt.Println("Aplikasi Parkir Server sedang on air")

	lis, error := net.Listen("tcp", "0.0.0.0:50051")

	if error != nil {
		log.Fatal("Tidak bisa menerima koneksi: %v", error)
	}

	s := grpc.NewServer()

	parkirpb.RegisterParkirServiceServer(s, &server{})

	if error := s.Serve(lis); error != nil {
		log.Fatalf("failed to serve: %v", error)
	}

}



func(*server) 	KeluarParkirProto(req *parkirpb.KeluarParkirRequest, stream parkirpb.ParkirService_KeluarParkirProtoServer) error {
	return nil
}



func (*server) 	MasukParkirProto(req *parkirpb.MasukParkirRequest,stream parkirpb.ParkirService_MasukParkirProtoServer) error{
	fmt.Println("Ada inputan dari Client terkait masuk parkir service dengan format: ", req)
	jenisKendaraan := req.GetMasukParkir().GetJenisKendaraan()
	platNomor := req.GetMasukParkir().GetPelatNomor()

	var parker PenggunaParkir
	generatedIdParkir := int64(idParkirGenerator())

	if generatedIdParkir != 0 {
		parker.idParkir = generatedIdParkir
		parker.jenisKendaraan = jenisKendaraan
		parker.platNomor = platNomor
		parker.waktuMasuk = time.Now().Second()

		var penggunaParkirArray PenggunaParkirArray
		penggunaParkirArray.penggunaParkirArray = append(penggunaParkirArray.penggunaParkirArray, parker)

		for i := 0; i < 10; i++ {
			res := &parkirpb.MasukParkirResponse {
				GeneratedId: generatedIdParkir,
			}
			stream.Send(res)

			time.Sleep(1000 * time.Millisecond)
		}
	}


	return nil
}

func idParkirGenerator() int {
	var flag bool = false
	var idParkirGenerated int
	for !flag {
		idParkirGenerated = rand.Intn(100)
		for i := 0; i < len(penggunaParkirArray); i++ {
			var person PenggunaParkir = penggunaParkirArray[i]
			if person.idParkir != int64(idParkirGenerated) {
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
		if person.idParkir != int64(idParkirMasuk) {
			penggunaParkirArray = append(penggunaParkirArray, penggunaParkir)
			return "Sudah masuk"
			break
		} else {
			fmt.Println("Parkir udah penuh atau id Parkir sama")
			break
		}
	}
	return "sudah penuh"
}

func parkirKeluar(idParkirKeluar int)  {
	for i := 0; i < len(penggunaParkirArray); i++ {
		var penggunaParkir1 PenggunaParkir = penggunaParkirArray[i]
		if penggunaParkir1.idParkir == int64(idParkirKeluar) {
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