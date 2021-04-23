package server

var UsersList []User

var LocationList []Location

type User struct {
	Id       int      `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Location Location `json:"location"`
	Groups   []string `json:"Groups"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
