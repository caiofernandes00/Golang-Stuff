package streaming

import (
	"testing"
	"time"
)

func Test_Listener(t *testing.T) {
	go func() {
		time.Sleep(4 * time.Second)
		sendFile(1000)
	}()
	server := &FileServer{}
	server.start()
}
