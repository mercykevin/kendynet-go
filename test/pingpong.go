package main

import(
	"net"
	"fmt"
)


func handleclient(Conn net.Conn){
	recvbuf := make([]byte,16000)
	for{
		_,err := Conn.Read(recvbuf)
		if err != nil {
			Conn.Close()
			return
		}
		_,err = Conn.Write(recvbuf)
		if err != nil {
			Conn.Close()
			return
		}
	}
}

func main(){
	service := ":8010"
	tcpAddr,err := net.ResolveTCPAddr("tcp4", service)
	if err != nil{
		fmt.Printf("ResolveTCPAddr")
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil{
		fmt.Printf("ListenTCP")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleclient(conn)
	}
}
