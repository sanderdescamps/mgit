package repo

import (
	"fmt"
	"io"
	"net"
	"time"
)

func oneOf(a string, list []string) bool {
	for _, i := range list {
		if a == i {
			return true
		}
	}
	return false
}

func tcpConnectionTest(host string, port int) error {
	endpoint := fmt.Sprintf("%s:%d", host, port)
	conn, _ := net.Dial("tcp", endpoint)

	err := conn.(*net.TCPConn).SetKeepAlive(true)
	if err != nil {
		return fmt.Errorf("failed to set KeepAlive")
	}

	err = conn.(*net.TCPConn).SetKeepAlivePeriod(30 * time.Second)
	if err != nil {
		return fmt.Errorf("failed to set KeepAlive period")
	}
	notify := make(chan error)

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				notify <- err
				if io.EOF == err {
					return
				}
			}
			if n > 0 {
				fmt.Println("unexpected data: %v", buf[:n])
			}
		}
	}()

	for {
		select {
		case err := <-notify:
			fmt.Println("connection dropped message", err)
			break
		case <-time.After(time.Second * 1):
			fmt.Println("timeout 1, still alive")
		}
	}
}
