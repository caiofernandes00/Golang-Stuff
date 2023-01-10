package streaming

import (
	"crypto/rand"
	"fmt"
	"io"
	"net"
)

func sendFile(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}

	n, err := conn.Write(file)
	if err != nil {
		return err
	}

	fmt.Printf("written %d bytes over the network\n", n)
	return nil
}
