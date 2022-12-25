package mutexes

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_mutex(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	example()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = stdOut

	expected := `$34320.00`

	if !strings.Contains(string(out), expected) {
		t.Errorf("Expected %s, got %s", expected, out)
	}
}
