package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/Taurin190/geofrontdb/internal/config"
	log "github.com/sirupsen/logrus"
)

const (
	TCP_PROTOCOL = "tcp"
)

func init() {
	flag.Parse()
}

func main() {
	log.Info("Starting Geofront Server...")
	config := config.NewServerConfig()

	tcpAddress, err := net.ResolveTCPAddr(TCP_PROTOCOL, fmt.Sprintf("localhost:%s", config.Port))
	if err != nil {
		log.Fatalf("Error resolving TCP address: %+v", err)
	}

	listener, err := net.ListenTCP(TCP_PROTOCOL, tcpAddress)
	if err != nil {
		log.Fatalf("Error listening: %+v", err)
	}
	log.Infof(fmt.Sprintf("[Server][Parent] *net.TCPListener Address: %s", listener.Addr().String()))

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			log.Warnf("Error accepting: %+v", err)
			continue
		}
		go handler(clientConn)
	}
}

func handler(conn net.Conn) {
	pid := os.Getpid()

	la := conn.LocalAddr().String()
	ra := conn.RemoteAddr().String()
	connInfo := fmt.Sprintf("[SERVER] pid:%d, localAddress:%s, remoteAddress:%s", pid, la, ra)
	log.Println(connInfo)

	for {
		request := make([]byte, 4096)

		// クライアント X から通信が来るまで待機する
		readLen, err := conn.Read(request)
		if err != nil {
			log.Printf("Error reading: %s", err.Error())
			break
		}

		// クライアントが接続を切った
		if readLen == 0 {
			break
		}

		conn.Write([]byte("[From][Server] Hello! Your message is " + string(request)))
		log.Printf("%s sent to client message: %s", connInfo, string(request))
	}

	conn.Close()
}
