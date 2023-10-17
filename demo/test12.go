package main

type User struct {
	Name string
	Age  int
	AA   string
	BB   string
}

func main() {
	up(User{
		Name: "name",
		Age:  10,
		AA:   "",
		BB:   "",
	}, "")
}
func up(a User, o string) string {
	if a.Name != "" {
		o = "111"
	}
	if a.Age == 0 {
		o = "222"
	}
	if a.AA == "c" {
		o = "333"
	}
	if a.BB == "d" {
		o = "444"
	}
	return o
}
