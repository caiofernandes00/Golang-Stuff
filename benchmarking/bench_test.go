package benchmarking

import (
	"bytes"
	"io"
	"testing"
)

func Benchmark_Slices(b *testing.B) {
	n := 100

	b.Run("NoAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := []int{}
			for j := 0; j < n; j++ {
				s = append(s, j)
			}
		}
	})

	b.Run("Alloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := make([]int, n)
			for j := 0; j < n; j++ {
				s[j] = j
			}
		}
	})
}

func writeToBuffer(msg []byte) {
	buf := new(bytes.Buffer)
	buf.Write(msg)
}

func writeToBufferWithAllocation(w io.Writer, msg []byte) {
	w.Write(msg)
}

func Benchmark_WriteToBuffer(b *testing.B) {
	msg := []byte("Hello World")

	b.Run("NoAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < 100; j++ {
				writeToBuffer(msg)
			}
		}
	})

	b.Run("Alloc", func(b *testing.B) {
		buf := new(bytes.Buffer)
		for i := 0; i < b.N; i++ {
			for j := 0; j < 100; j++ {
				writeToBufferWithAllocation(buf, msg)
				buf.Reset()
			}
		}
	})
}
