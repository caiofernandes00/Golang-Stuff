package slices

type MySlice []int

func (s MySlice) Remove(index int) []int {
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}

func (s MySlice) RemoveAndKeepOrder(index int) []int {
	return append(s[:index], s[index+1:]...)
}
