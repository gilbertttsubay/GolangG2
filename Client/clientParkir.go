package main



import (
	"TugasParkir/gRPC"
	"log"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main(){

	cc, error := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if error != nil {
		log.Fatalf("Could not connect becasause: ", error)
	}

	

	c := tugasparkirpb.NewParkirServiceClient(cc)
	fmt.Printf("Created client: %f", c)

	menu()
	var flag bool = false

	for !flag {
		fmt.Println("Masukkan angka")
		var angkaSwitch int
		fmt.Scan(&angkaSwitch)
		
		switch angkaSwitch {
		case 1:
			masukParkir(c)
			break
		case 2:
			keluarParkir(c)
			break
		case 3:
			fmt.Println("Keluar menu")
			defer cc.Close()
			flag = true
			break
		}
	}
}

func menu(){
	fmt.Println("Menu Parkir")
	fmt.Println("1. Masuk Parkir")
	fmt.Println("2. Keluar Parkir")
	fmt.Println("3. Keluar menu")
}

func masukParkir(c tugasparkirpb.ParkirServiceClient){
	var (
		jenisKendaraan string;
		platNomor string;
	)
	fmt.Println("Masuk Parkir")
	fmt.Println("Masukkan jenis Kendaraan")
	fmt.Scan(&jenisKendaraan)
	
	fmt.Println("Masukkan plat nomor")
	fmt.Scan(&platNomor)

	req := &tugasparkirpb.MasukParkirRequest{
		MasukParkir: &tugasparkirpb.MasukParkir{
			JenisKendaraan: jenisKendaraan,
			PelatNomor:  platNomor,
		},
	}
	res, err := c.MasukParkir(context.Background(), req)
	if err != nil {
		log.Fatalf("error response dari Server Parkir: %v", err)
	}
	log.Printf("Id Parkir Anda Adalah: %v", res.GeneratedId)
}

func keluarParkir(c tugasparkirpb.ParkirServiceClient){
	var(
		idParkirDariServer int64
	)

	fmt.Println("Masukkan id Parkir")
	fmt.Scan(&idParkirDariServer)

	req:= &tugasparkirpb.KeluarParkirRequest{
		KeluarParkir: &tugasparkirpb.KeluarParkir{
			IdKeluar: idParkirDariServer,
		},
	}

	res, err := c.KeluarParkir(context.Background(), req)
	if err != nil {
		log.Fatalf("error response dari Server Parkir: &v", err)
	}
	log.Printf("Total biaya Anda adalah: %v", res.TotalBiaya)
}