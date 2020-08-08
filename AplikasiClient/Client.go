package main
import(
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"log"
	"io/ioutil"
)


type kendaraanMasuk struct{
	jenisKendaraan string
	platNomor string
}

type idParkir struct {
	IdParkir int `json : "IdParkir`
}
type totalBiaya struct {
	TotalBiaya int64 `json: "TotalBiaya`
}

var baseURL = "http://localhost:8080"

func main(){
	http.HandleFunc("/masukparkir", parkirMasuk)

	fmt.Println("Client is running...")

	http.ListenAndServe(":8888", nil)


}

func menu(){
	fmt.Println("Selamat datang di menu parkir G2 Academy")
	fmt.Println("1. Masuk parkir")
	fmt.Println("2. Keluar parkir")
	fmt.Println("3. Keluar menu")
}

func parkirMasuk(res http.ResponseWriter, req *http.Request) {
	var client = &http.Client{}

	res.Header().Set("Content-Type", "application/json")

	//request ke postman (input jenisKendaraan dan plat nomor)
	if req.Method == "POST" {
		res.WriteHeader(200)
		//response dari postman
		responseDariPostman,error:= ioutil.ReadAll(req.Body)
		if error != nil {
			log.Fatalf("Tidak menerima hasil inputan dari postman dikarenakan: ", error)
		}

		//meminta request ke Server melalui url yang menyambungkannya
		requestKeServer,error := http.NewRequest("POST","http://localhost:8080/id_parkir", bytes.NewBuffer(responseDariPostman))
		if error != nil {
			log.Fatalf("Tidak bisa mengirim request Post untuk masuk Parkir ke Server dikarenakan: ", error)
		}
		requestKeServer.Header.Set("Content-Type", "application/json")

		//response dari server
		responseDariServer, error := client.Do(requestKeServer)
		if error != nil {
			log.Fatalf("Tidak bisa menerima response dari Server terkait inputan untuk masuk parkir dikarenakan: ", error)
		}
		defer responseDariServer.Body.Close()

		//mengolah data response yang didapat dari server
		dataResponse, error := ioutil.ReadAll(responseDariServer.Body)
		if error != nil {
			log.Fatalf("Tidak bisa mengolah data dari server karena: ", error)
		}

		

		//menampilkan hasil response dari server ke postman yang di bawah
		res.Write([]byte(string(dataResponse)))
	}
}

func parkirKeluar(idParkir int) int64{
	var baseURL = "http://localhost:8080"
	var client= &http.Client{}
	var totalBiaya totalBiaya


	var param = url.Values{}
	param.Set("id_Parkir", string(idParkir))
	var payload = bytes.NewBufferString(param.Encode())
	

	request, err := http.NewRequest("POST", baseURL+"/keluar", payload)
	if err != nil {
		log.Fatalf("Tidak bisa konek ke post karena ", err)
	}

	//response dari server

	response, error := client.Do(request)
	if error != nil {
		log.Fatalf("Gagal menerima response keluar parkir dari Server karena ", error)
	}

	defer response.Body.Close()

	error = json.NewDecoder(response.Body).Decode(&totalBiaya)

	if error != nil {
		log.Fatalf("Gagal menerima total biaya dari server dikarenakan: ", error)
	}

	return totalBiaya.TotalBiaya
}