package comms

import (
	"log"
	"net"
)

func Listen(host string, port int) {
	ip := net.ParseIP(host)
	addr := &net.TCPAddr{IP: ip, Port: port}
	log.Printf("listening on %s", addr)
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		log.Printf("Connect: %s...\n", conn.RemoteAddr())
		//output.addWriteTo(conn)
	}

}
