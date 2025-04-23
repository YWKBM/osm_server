package entities

type Zone struct {
	Id        int
	Name      string
	Providers []Provider
	Geo       string
}
