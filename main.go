package main

type User struct {
	UUID     int
	surname  string
	name     string
	patronym string
	age      int
}

type Student struct {
	student   User
	code      string
	course    int
	group     int
	faculties string
}

type Teacher struct {
	teacher      User
	cafedra      string
	faculties    string
	science_rang string
}

func main() {

}
