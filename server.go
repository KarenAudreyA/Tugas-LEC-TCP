package main 
import(
	"net"
	"fmt"
)

func main(){

	listener, err := net.Listen("tcp","localhost:1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	// defer -> akan menjalankan code disampingnya ketika function yg ada defernya sudah selesai dijalankan
	defer listener.Close()

	// buat loop infinite spaya selalu bs nerima req dr client
	for{
		// acc dari req client
		// conn digunakan utk mnerima pesan yg dikirim sm client dan bisa tau isinya dgn di print
		// conn, err := net.Accept()
		_ , err := listener.Accept()

		if err != nil{
			fmt.Println(err)
			return
		}

		// fungsi go routine ini membuat server tetap bisa menerima req dari client smbil mengeksekusi req client yg skrng
		go func(){
			fmt.Println("Accept")
			
		}()
		// () di bawah itu utk passing datanya ke parameter diatas
	}

}