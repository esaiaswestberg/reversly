package server

import (
	"io"
	"log"
	"net"
)

func ListenAndProxy(listenAddr string, proxyAddr string) error {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go func() {
			proxy, err := net.Dial("tcp", proxyAddr)
			if err != nil {
				return
			}

			go func() {
				_, err := io.Copy(conn, proxy)
				if err != nil {
					log.Println(err)
				}
			}()

			go func() {
				_, err := io.Copy(proxy, conn)
				if err != nil {
					log.Println(err)
				}
			}()
		}()
	}
}
