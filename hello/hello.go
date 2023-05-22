package hello

import (
	"time"

	"github.com/microsoft/vscode-remote-try-go/geoloc"
)

type Params struct {
	Name string
	Ip   string
}

// User user type
type User struct {
	ID   int64
	Name string
	Addr *Address
}

// Address address type
type Address struct {
	Country string
	ZIP     int
	LatLng  [2]float64
}

func Register(params Params) User {
	name := params.Name
	country := geoloc.Lookup_struct(params.Ip)
	return User{
		ID:   time.Now().UnixNano(),
		Name: name,
		Addr: &Address{
			Country: country,
			ZIP:     1488,
			LatLng:  [2]float64{0.0, 9.0},
		}}
}

// Hello writes a welcome string
func Hello(user User) string {
	return "Hello, " + user.Name + "! I see you are somewhere from " + user.Addr.Country
}
