package slices

var users = []string{}

func addUsers1(users []string) {
	for _, user := range users {
		users = append(users, user)
	}
}

func addUsers2(users ...string) {
	for _, user := range users {
		users = append(users, user)
	}
}

func addUsers3(user ...string) {
	users = append(users, user...)
}
