package slices

import "testing"

func Test_AddUsersSlices1(t *testing.T) {
	users := []string{"a", "b"}
	addUsers1(users)
	if len(users) != 2 {
		t.Errorf("incorrect length of slice; expected 2 but got %d", len(users))
	}
}

func Test_AddUsersSlices2(t *testing.T) {
	users := []string{"a", "b"}
	addUsers2("a", "b")
	if len(users) != 2 {
		t.Errorf("incorrect length of slice; expected 2 but got %d", len(users))
	}
}

func Test_AddUsersSlices3(t *testing.T) {
	users := []string{"a", "b"}
	addUsers3("a", "b")
	if len(users) != 2 {
		t.Errorf("incorrect length of slice; expected 2 but got %d", len(users))
	}
}
