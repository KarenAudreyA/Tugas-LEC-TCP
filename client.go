package main
import (
	"net"
	"fmt"
)

func main(){

	// buat func yg bisa ngirim ke server
	// local host -> IP 127.0.0.1
	// cara 1
		// net.Dial("tcp","localhost:1234")
	// cara 2
	dial,err :=	net.Dial("tcp","127.0.0.1:1234")
	if err != nil{
		fmt.Println(err)
		return

	}
	defer dial.Close()	
	


}