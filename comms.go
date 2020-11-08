package main

import "net"

func listenForClient() {
	addr := &net.TCPAddr{Port: 9998}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
}
