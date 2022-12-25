package safe_map

import (
	"fmt"
	"testing"
)

func Test_SafeMap(t *testing.T) {
	safeMap := New[int, int]()

	for i := 0; i < 10; i++ {
		go func(i int) {
			safeMap.Insert(i, i*2)
			value, _ := safeMap.Get(1)
			fmt.Println(value)
		}(i)
	}
}
