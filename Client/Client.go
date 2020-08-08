package main

import (
	"io"
	"log"
	"google.golang.org/grpc"
	"Parkir_gRPC_Server_Streaming_rpc/Proto"
	"fmt"
	"context"
)



func main (){
	cc, error := grpc.Dial("localhost:50051",grpc.WithInsecure())

	if error != nil {
		log.Fatalf("Tidak bisa konek ke localhost 50051 karena ada: ", error)
	}

	c := parkirpb.NewParkirServiceClient(cc)
	fmt.Println("Aplikasi Client Parkir sedang on Air")

	var flag bool
	for !flag {
		fmt.Println("Selamat datang di menu Parkir G2 Academy")
		fmt.Println("Silahkan masukkan angka")
		var angkaSwitch int

		fmt.Scan(&angkaSwitch)
		switch angkaSwitch {
		case 0:
			menu()
			defer cc.Close()
			break
		case 1:
			masukParkir(c)
			defer cc.Close()
			break;
		}
	}

}


func menu(){
	fmt.Println("Menu Parkir")
	fmt.Println("0. Menu")
	fmt.Println("1. Masuk Parkir")
	fmt.Println("2. Keluar Parkir")
	fmt.Println("3. Keluar menu")
}

func masukParkir(c parkirpb.ParkirServiceClient) {
	var (
		jenisKendaraan string;
		platNomor string;
	)

	fmt.Println("Masuk Parkir")

	fmt.Println("Masukkan jenis Kendaraan")
	fmt.Scan(&jenisKendaraan)

	fmt.Println("Masukkan plat nomor")
	fmt.Scan(&platNomor)

	request := &parkirpb.MasukParkirRequest{
		MasukParkir: &parkirpb.MasukParkir {
			JenisKendaraan: jenisKendaraan,
			PelatNomor: platNomor,
		},
	}

	resStream, error:= c.MasukParkirProto(context.Background(), request)
	if error != nil {
		log.Fatalf("Tidak bisa mendapatkan hasil streaming dari server dikarenakan: ", error)
	}
	for {
		msg,error := resStream.Recv()
		if error != nil {
			log.Fatal("Tidak bisa mendapatkan pesan dari streaming karena: ", error)
			break
		}
		if error == io.EOF {
			break
		}

		fmt.Println("Id parkir Anda adalah: ", msg.GetGeneratedId())
	}

}