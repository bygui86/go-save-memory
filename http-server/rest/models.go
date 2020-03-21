package rest

type User struct {
	Name       string
	Surname    string
	Age        int
	BestMovies map[string]string // key: genre - value: title
	Hobbies    []string
}
